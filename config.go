package main

const (
	KafkaBootstrapServers = "localhost:39092"
	ServiceAddress        = "localhost:8888"
	TopicName             = "email-updated"
	ConsumerGroupID       = "email-verifier"
	CSRInstanceBase       = "https://buf-gophercon.buf.dev/integrations/confluent/"

	CSRInstanceName = "<CSR INSTANCE NAME>"
	BSRUser         = "gophercon-csr-demo"
	BSRToken        = "<BSR TOKEN>"
)
