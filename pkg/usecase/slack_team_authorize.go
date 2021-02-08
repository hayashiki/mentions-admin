package usecase

import (
	"context"
	"fmt"
	"github.com/hayashiki/go-pkg/slack/auth"
	"github.com/hayashiki/mentions/pkg/model"
	"github.com/hayashiki/mentions/pkg/repository"
	"time"
)

type TeamAuth interface {
	Do(ctx context.Context, input *TeamAuthInput) error
}

type teamAuth struct {
	token    auth.Token
	teamRepo repository.TeamRepository
}

func NewTeamAuth(token auth.Token, teamRepo repository.TeamRepository) TeamAuth {
	return &teamAuth{
		token:    token,
		teamRepo: teamRepo,
	}
}

type TeamAuthInput struct {
	Code string
}

func (uc *teamAuth) Do(ctx context.Context, input *TeamAuthInput) error {
	authResp, err := uc.token.GetAccessToken(input.Code)
	if err != nil {
		return fmt.Errorf("fail to get token %v", err)
	}

	team := &model.Team{
		ID:        authResp.Team.ID,
		Name:      authResp.Team.Name,
		Token:     authResp.AccessToken,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return uc.teamRepo.Put(ctx, team)
}
