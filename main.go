package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := VerifyTopic(ctx); err != nil {
		log.Fatal(err)
	}

	producer := must(NewProducer())
	defer producer.Close()

	consumer := must(NewConsumer())
	defer consumer.Close()

	csrClient := must(NewClient())
	store := must(NewStore(csrClient, producer))

	enabled := &atomic.Bool{}
	enabled.Store(true)
	verifier := must(NewVerifier(store, csrClient, consumer, enabled))
	service := NewService(store, enabled)

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error { return service.Run(ctx) })
	grp.Go(func() error { return verifier.Run(ctx) })

	if err := grp.Wait(); err != nil {
		log.Fatal(err)
	}
}

func must[T any](value T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return value
}
