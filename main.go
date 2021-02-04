package main

import (
	"fmt"
	"github.com/hayashiki/mentions-admin/pkg/app"
	"github.com/hayashiki/mentions-admin/pkg/config"
	"log"
	"net/http"
	"os"
)

//go:generate go run go.pyspa.org/brbundle/cmd/brbundle embedded -f web/dist
func main() {
	config := config.MustReadConfigFromEnv()
	app, err := app.NewApp(config)

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), app.Handler()); err != nil {
		log.Fatal(err)
	}
}
