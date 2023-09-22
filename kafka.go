package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/protobuf/proto"
)

func kafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": KafkaBootstrapServers,
	}
}

func NewProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(kafkaConfig())
}

func NewConsumer() (*kafka.Consumer, error) {
	cfg := kafkaConfig()
	if err := cfg.SetKey("group.id", ConsumerGroupID); err != nil {
		return nil, err
	}
	return kafka.NewConsumer(cfg)
}

func VerifyTopic(ctx context.Context) error {
	client, err := kafka.NewAdminClient(kafkaConfig())
	if err != nil {
		return err
	}
	defer client.Close()

	md, err := client.GetMetadata(proto.String(TopicName), false, 5*1000 /*ms*/)
	if err != nil {
		return err
	}

	if topic, ok := md.Topics[TopicName]; ok && topic.Error.Code() == 0 {
		// topic exists
		return nil
	}

	res, err := client.CreateTopics(ctx, []kafka.TopicSpecification{{
		Topic:         TopicName,
		NumPartitions: 1,
	}})
	if err != nil {
		return err
	}
	if e := res[0].Error; e.Code() != 0 && e.Code() != kafka.ErrTopicAlreadyExists {
		return fmt.Errorf("failed to create topic: %s", e.Error())
	}
	return nil
}
