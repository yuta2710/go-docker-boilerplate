package main

import (
	"fmt"

	"github.com/yuta_2710/go-clean-arc-reviews/config"
	"github.com/yuta_2710/go-clean-arc-reviews/database"
	"github.com/yuta_2710/go-clean-arc-reviews/server"
)

// "github.com/yuta_2710/go-clean-arc-reviews/config"
// "github.com/yuta_2710/go-clean-arc-reviews/database"
// "github.com/yuta_2710/go-clean-arc-reviews/server"

func main() {
	fmt.Println("Hello Docker")
	conf := config.GetConfig()
	postgres := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, postgres).Start()
}