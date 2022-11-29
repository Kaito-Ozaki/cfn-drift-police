package slackcli

import (
	"fmt"

	"github.com/slack-go/slack"
)

type InMemorySlackClient struct {
	svc *slack.Client
}

func NewInMemorySlackClient(token string) SlackClient {
	fmt.Println(token)
	return InMemorySlackClient{}
}

func (cli InMemorySlackClient) PostMessage(channelName string, options ...slack.MsgOption) error {
	return nil
}
