package app

import (
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/hayashiki/mentions-admin/pkg/config"
	"github.com/hayashiki/mentions/pkg/repository"
	"go.pyspa.org/brbundle"
	"go.pyspa.org/brbundle/brchi"
	"log"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Printf("failed to process request: %s [%s]", err.message, err.err.Error())
		switch err {
		case errBadRequest:
			http.Error(w, err.message, http.StatusBadRequest)
		case errUnauthorized:
			http.Error(w, err.message, http.StatusUnauthorized)
		case errForbidden:
			http.Error(w, err.message, http.StatusForbidden)
		case errNotFound:
			http.Error(w, err.message, http.StatusNotFound)
		case errMethodNotAllowed:
			http.Error(w, err.message, http.StatusMethodNotAllowed)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type App struct {
	isDev    bool
	slack    config.Slack
	teamRepo repository.TeamRepository
}

func NewApp(config *config.Config, teamRepo repository.TeamRepository) (*App, error) {
	return &App{
		isDev:    config.IsDev,
		slack:    config.Slack,
		teamRepo: teamRepo,
	}, nil
}

func (app *App) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
	)
	//r.Use(cors.Handler(cors.Options{
	//	// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
	//	AllowedOrigins: []string{"*"},
	//	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: false,
	//	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//}))

	r.Method(http.MethodGet, "/health", appHandler(app.healthCheck))
	r.Method(http.MethodGet, "/slack/team/auth", appHandler(app.slackTeamAuthorize))
	r.Method(http.MethodGet, "/slack/team/callback", appHandler(app.slackTeamCallback))

	r.NotFound(brchi.Mount(brbundle.WebOption{
		SPAFallback: "index.html",
	}))
	return r
}

func (app *App) healthCheck(w http.ResponseWriter, r *http.Request) *appError {
	userEmail := r.Header.Get("X-Goog-Authenticated-User-Email")
	userId := r.Header.Get("X-Goog-Authenticated-User-ID")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request by: %s, %s\n", userEmail, userId)
	return nil
}
