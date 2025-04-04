package service

import (
	"github.com/DataDog/datadog-go/statsd"
	"log"
)

const (
	_serverAddress = "127.0.0.1"
	_namespace     = "demo-twelve.dev."
)

func NewDatadogClient() *statsd.Client {
	client, err := statsd.New("localhost:8125")
	if err != nil {
		log.Fatal(err)
	}

	client.Tags = append(client.Tags, "env:development")
	defer client.Flush()

	log.Println("Conexión a datadog realizada exitosamente")

	client.Namespace = _namespace
	return client
}
