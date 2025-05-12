// @title           alexKudryavtsev-web/default-service
// @version         1.0
// @description     default-service

// @BasePath  /api
// @schemes https http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't init config: %s", err)
	}

	app.Run(cfg)
}
