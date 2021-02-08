package app

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hayashiki/go-pkg/slack/auth"
	"github.com/hayashiki/mentions-admin/pkg/usecase"

	"net/http"
	"time"
)

var (
	UserScopes = []string{"identity.basic"}
	TeamScopes = []string{"chat:write", "users.profile:read", "users:read"}
)

func (app *App) slackTeamAuthorize(w http.ResponseWriter, r *http.Request) *appError {
	state := uuid.New().String()
	authz := auth.NewAuth(app.slack.ClientID, app.slack.RedirectURL, state, TeamScopes)

	http.SetCookie(w, &http.Cookie{
		Name:    "state",
		Value:   state,
		Expires: time.Now().Add(300 * time.Second),
		Secure:  true,
	})

	authz.Redirect(w, r)

	return nil
}

// slackTeamCallback
func (app *App) slackTeamCallback(w http.ResponseWriter, r *http.Request) *appError {
	resp := auth.ParseRequest(r)
	state, err := r.Cookie("state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return &appError{
			err:     fmt.Errorf("invalid cookie, err=%v", err),
			message: "invalid cookie",
		}

	}
	if resp.State != state.Value {
		return &appError{
			err:     fmt.Errorf("invalid state, err=%v", err),
			message: "invalid state",
		}

	}
	token := auth.NewToken(app.slack.ClientID, app.slack.SecretID, app.slack.RedirectURL)
	authz := usecase.NewTeamAuth(token, app.teamRepo)
	input := &usecase.TeamAuthInput{Code: resp.Code}
	if err := authz.Do(r.Context(), input); err != nil {
		return &appError{
			err:     fmt.Errorf("failed to authorize, err=%v", err),
			message: "invalid authorize",
		}
	}
	// TODO: リダイレクトする
	fmt.Fprintf(w, "team is: %s", "ok")
	return nil
}
