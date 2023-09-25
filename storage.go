package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"connectrpc.com/connect"
	emailsv1 "github.com/bufbuild/gophercon-2023-workshop/gen/emails/v1"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"google.golang.org/protobuf/proto"
)

type Storage struct {
	store      sync.Map
	serializer serde.Serializer
	producer   *kafka.Producer
}

func NewStore(csrClient schemaregistry.Client, producer *kafka.Producer) (*Storage, error) {
	serializer, err := NewSerializer(csrClient)
	if err != nil {
		return nil, err
	}
	return &Storage{
		serializer: serializer,
		producer:   producer,
	}, nil
}

func (s *Storage) GetEmail(userID uint64) (*emailsv1.UserEmail, error) {
	if val, exists := s.store.Load(userID); exists {
		return val.(*emailsv1.UserEmail), nil
	}
	return nil, connect.NewError(connect.CodeNotFound,
		fmt.Errorf("no email found for user %d", userID))
}

func (s *Storage) UpdateEmail(userID uint64, addr string) error {
	prev, hasPrev := s.store.Swap(userID, &emailsv1.UserEmail{
		UserId:       userID,
		EmailAddress: addr,
		Verified:     false,
	})

	event := &emailsv1.EmailUpdated{
		UserId:     userID,
		NewAddress: addr,
	}
	if hasPrev {
		event.OldAddress = prev.(*emailsv1.UserEmail).EmailAddress
	}

	s.emitEvent(event)
	return nil
}

func (s *Storage) emitEvent(event *emailsv1.EmailUpdated) {
	topicName := TopicName
	payload, err := s.serializer.Serialize(topicName, event)
	if err != nil {
		log.Printf("failed to serialize event: %v", err)
	}

	delivery := make(chan kafka.Event, 1)
	err = s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicName,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(strconv.FormatUint(event.GetUserId(), 10)),
		Value: payload,
	}, delivery)
	if err != nil {
		log.Printf("failed to produce message: %v", err)
	}

	log.Printf("published: %v", <-delivery)
}

func (s *Storage) VerifyEmail(userID uint64, addr string) error {
	val, exists := s.store.Load(userID)
	if !exists {
		return connect.NewError(connect.CodeNotFound,
			fmt.Errorf("no email found for user %d", userID))
	}

	record := val.(*emailsv1.UserEmail)
	if record.EmailAddress != addr {
		return connect.NewError(connect.CodeInvalidArgument,
			fmt.Errorf("verified address %q does not match user's address %q", addr, record.EmailAddress))
	}

	record = proto.Clone(record).(*emailsv1.UserEmail)
	record.Verified = true
	s.store.Store(userID, record)
	return nil
}
