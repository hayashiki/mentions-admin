package main

import (
	"fmt"
	"github.com/hayashiki/mentions-admin/pkg/app"
	"github.com/hayashiki/mentions-admin/pkg/config"
	"github.com/hayashiki/mentions/pkg/repository"
	"log"
	"net/http"
	"os"
)

//go:generate go run go.pyspa.org/brbundle/cmd/brbundle embedded -f web/dist
func main() {
	config := config.MustReadConfigFromEnv()
	teamRepo := repository.NewTeamRepository(repository.GetClient(config.GCPProject))

	app, err := app.NewApp(config, teamRepo)

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
