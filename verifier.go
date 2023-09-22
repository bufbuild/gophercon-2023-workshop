package main

import (
	"context"
	"log"

	emailsv1 "github.com/bufbuild/gophercon-2023-workshop/gen/emails/v1"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
)

type VerifierDaemon struct {
	store        *Storage
	consumer     *kafka.Consumer
	deserializer serde.Deserializer
}

func NewVerifier(
	store *Storage,
	csrClient schemaregistry.Client,
	consumer *kafka.Consumer,
) (*VerifierDaemon, error) {
	des, err := NewDeserializer(csrClient)
	if err != nil {
		return nil, err
	}
	return &VerifierDaemon{
		store:        store,
		consumer:     consumer,
		deserializer: des,
	}, nil
}

func (v *VerifierDaemon) Run(ctx context.Context) error {
	err := v.consumer.SubscribeTopics([]string{TopicName}, nil)
	if err != nil {
		return err
	}
	defer v.consumer.Unsubscribe()
	defer v.consumer.Commit()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err = v.consume(); err != nil {
				return err
			}
		}
	}
}

func (v *VerifierDaemon) consume() error {
	event := v.consumer.Poll(1000 /*ms*/)
	if event == nil {
		return nil
	}

	switch e := event.(type) {
	case kafka.Error:
		return e
	case *kafka.Message:
		msg, err := v.deserializer.Deserialize(*e.TopicPartition.Topic, e.Value)
		if err != nil {
			return err
		}
		updated, ok := msg.(*emailsv1.EmailUpdated)
		if !ok {
			log.Printf("unexpected message received: %v", msg)
		} else if err = v.store.VerifyEmail(updated.GetUserId(), updated.GetNewAddress()); err != nil {
			log.Printf("failed to verify email: %v", err)
		} else {
			log.Printf("verified email update for %d: %s -> %s",
				updated.GetUserId(), updated.GetOldAddress(), updated.GetNewAddress())
		}
		return nil
	default:
		return nil
	}
}
