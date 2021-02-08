package usecase

import (
	"context"
	"fmt"
	"github.com/hayashiki/go-pkg/slack/auth"
	"github.com/hayashiki/mentions/pkg/model"
	"github.com/hayashiki/mentions/pkg/repository"
	"github.com/hayashiki/mentions-admin/pkg/slack"
	"time"
)

type TeamAuth interface {
	Do(ctx context.Context, input *TeamAuthInput) error
}

type teamAuth struct {
	token    auth.Token
	teamRepo repository.TeamRepository
	userRepo repository.UserRepository
}

func NewTeamAuth(token auth.Token, teamRepo repository.TeamRepository, userRepo repository.UserRepository) TeamAuth {
	return &teamAuth{
		token:    token,
		teamRepo: teamRepo,
		userRepo: userRepo,
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

	if err := uc.teamRepo.Put(ctx, team); err != nil {
		return fmt.Errorf("failed to put user team: %v, errpr: %w", team, err)
	}

	slackSvc := slack.NewClient(slack.New(team.Token))
	users, err := slackSvc.GetUsers()
	// TODO: MultiPut
	// TODO: team.IDをわたす
	for _, user := range users {
		if err := uc.userRepo.Put(ctx, team, user); err != nil {
			// TODO: errを配列で格納する
			fmt.Errorf("failed to put user user:%v, errpr: %w", user, err)
		}
	}
	return nil
}
