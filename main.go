package main

import (
	"flag"
	tgClient "github.com/gymnasy55/adviser-bot/clients/telegram"
	"github.com/gymnasy55/adviser-bot/consumer/event-consumer"
	"github.com/gymnasy55/adviser-bot/events/telegram"
	"github.com/gymnasy55/adviser-bot/storage/files"
	"log"
	"path/filepath"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

var (
	storagePath = filepath.Join(".", "files_storage")
)

func main() {
	client := tgClient.New(tgBotHost, mustToken())
	storage := files.New(storagePath)
	eventsProcessor := telegram.New(client, storage)
	log.Print("service started...")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("bot-token", "", "token for access telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
}
