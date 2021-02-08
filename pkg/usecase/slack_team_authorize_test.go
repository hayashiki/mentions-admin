package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/hayashiki/go-pkg/slack/auth"
	"github.com/hayashiki/mentions/pkg/model"
	"github.com/hayashiki/mentions/pkg/repository"
	mock_repo "github.com/hayashiki/mentions/pkg/repository/mock_repo"
	"testing"
	"time"
)

func NewMockToken() auth.Token {
	return MockToken{}
}

type MockToken struct {
	Error error
}

func (t MockToken) GetAccessToken(code string) (authResp auth.TokenResponse, err error) {
	if t.Error != nil {
		return auth.TokenResponse{}, t.Error
	}

	return auth.TokenResponse{
		Ok:          true,
		AppID:       "A0148GGR6AF",
		AuthedUser:  auth.AuthedUser{
			ID: "UD7AKTEFV",

		},
		Scope:       "chat:write,incoming-webhook,users.profile:read,users:read",
		TokenType:   "bot",
		AccessToken: "xoxb-1234567890",
		BotUserID:   "U014PE5ESUC",
		Team:        auth.TeamInfo{
			ID:"T0CG2CUSY",
			Name: "dummy",
		},
		Enterprise:  "",
		Error:       "",
	}, nil
}

func Test_teamAuth_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	teamMockRepo := mock_repo.NewMockTeamRepository(ctrl)

	teamOk := &model.Team{
		ID:        "T0CG2CUSY",
		Name:      "dummy",
		Token:     "xoxb-1234567890",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	type fields struct {
		token    auth.Token
		teamRepo repository.TeamRepository
	}
	type args struct {
		ctx   context.Context
		input *TeamAuthInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		respTeam *model.Team
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				token:    NewMockToken(),
				teamRepo: teamMockRepo,
			},
			args: args{
				ctx:   context.Background(),
				input: &TeamAuthInput{
					Code: "dummy",
				},
			},
			wantErr: false,
			respTeam: teamOk,
		},
	}
	for _, tt := range tests {
		// TODO: gomock.Anyをつかわずに、timeモックをさしこむ
		teamMockRepo.EXPECT().Put(tt.args.ctx, gomock.Any()).Return(nil)
		t.Run(tt.name, func(t *testing.T) {
			uc := &teamAuth{
				token:    tt.fields.token,
				teamRepo: tt.fields.teamRepo,
			}
			if err := uc.Do(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
