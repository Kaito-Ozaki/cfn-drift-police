package slackcli

import (
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
)

type DefaultSlackClient struct {
	svc *slack.Client
}

func NewDefaultSlackClient(token string) SlackClient {
	return DefaultSlackClient{
		svc: slack.New(token),
	}
}

func (cli DefaultSlackClient) PostMessage(channelName string, options ...slack.MsgOption) error {
	_, _, err := cli.svc.PostMessage(channelName, options...)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
