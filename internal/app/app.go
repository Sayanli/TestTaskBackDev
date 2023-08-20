package app

import (
	"fmt"

	"github.com/Sayanli/TestTaskBackDev/internal/config"
	"github.com/Sayanli/TestTaskBackDev/pkg/database/mongodb"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	mongoClient, err := mongodb.NewClient(cfg.MongoDB.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	fmt.Println(mongoClient)
}
