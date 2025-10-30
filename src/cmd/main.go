package main

import (
	"fmt"
	"tasty-discoveries/src/config"
	"tasty-discoveries/src/internal/repository/opensearch"
)

func main() {
	config, _ := config.Load()

	_, err := opensearch.NewClient(
		opensearch.Config{
			Host:     config.DbHost,
			Port:     config.DbPort,
			Password: config.DbPassword,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}
