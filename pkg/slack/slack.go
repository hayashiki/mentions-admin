package slack

import (
	"github.com/hayashiki/mentions/pkg/model"
	"github.com/slack-go/slack"
	"regexp"
)

var r = regexp.MustCompile(`@[a-zA-Z0-9_\-\.]+`)

type Client interface {
	GetUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}

type client struct {
	bot *slack.Client
}

func NewClient(cli *slack.Client) Client {
	return &client{bot: cli}
}

func New(token string) *slack.Client {
	return slack.New(token)
}

// GetUsers
func (c *client) GetUsers() ([]*model.User, error) {
	slackUsers, err := c.bot.GetUsers()
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, user := range slackUsers {

		if user.IsBot {
			continue
		}
		if user.Deleted {
			continue
		}
		if user.IsInvitedUser {
			continue
		}

		name := user.Profile.DisplayName
		if name == "" {
			name = user.Name
		}

		//IsRestricted、IsAdmin、IsOwnerがほしい

		users = append(users, &model.User{
			ID:     user.ID,
			Name:   name,
			Avatar: user.Profile.Image192,
			TeamID: user.TeamID,
		})
	}

	return users, nil
}

func (c *client) GetUser(id string) (*model.User, error) {
	user, err := c.bot.GetUserInfo(id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:     user.ID,
		Name:   user.Profile.DisplayName,
		Avatar: user.Profile.Image192,
	}, nil
}

type MessageResponse struct {
	Channel   string
	Timestamp string
}
