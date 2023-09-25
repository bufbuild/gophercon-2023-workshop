package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func NewClient() (schemaregistry.Client, error) {
	cfg := schemaregistry.NewConfigWithAuthentication(
		CSRInstanceBase+CSRInstanceName, BSRUser, BSRToken)
	return schemaregistry.NewClient(cfg)
}

func NewSerializer(client schemaregistry.Client) (serde.Serializer, error) {
	return protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
}

func NewDeserializer(client schemaregistry.Client) (serde.Deserializer, error) {
	des, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		return nil, err
	}
	des.ProtoRegistry = protoregistry.GlobalTypes
	return des, nil
}
