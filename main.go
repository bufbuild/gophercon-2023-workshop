package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := VerifyTopic(ctx); err != nil {
		log.Fatal(err)
	}

	csrClient, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	producer, err := NewProducer()
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	consumer, err := NewConsumer()
	if err != nil {
		log.Fatal(err)
	}

	store, err := NewStore(csrClient, producer)
	if err != nil {
		log.Fatal(err)
	}

	verifier, err := NewVerifier(store, csrClient, consumer)
	if err != nil {
		log.Fatal(err)
	}

	service := NewService(store)

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error { return service.Run(ctx) })
	grp.Go(func() error { return verifier.Run(ctx) })

	if err := grp.Wait(); err != nil {
		log.Fatal(err)
	}
}
