package config

import (
	"flag"
	"log"
)

type Config struct {
	TgBotToken            string
	MongoConnectionString string
}

func MustLoad() Config {
	tgBotTokenToken := flag.String(
		"tg-bot-token",
		"7073460229:AAG4nVGKjjXz2if1t2HccLr9IE682jI8_fk",
		"token for access to telegram bot",
	)

	mongoConnectionString := flag.String(
		"mongo-connection-string",
		"mongodb://root:example@localhost:27017",
		"connection string for MongoDB",
	)

	flag.Parse()

	if *tgBotTokenToken == "" {
		log.Fatal("токен не указан")
	}
	if *mongoConnectionString == "" {
		log.Fatal("mongo connection-строка не указана")
	}

	return Config{
		TgBotToken:            *tgBotTokenToken,
		MongoConnectionString: *mongoConnectionString,
	}
}
