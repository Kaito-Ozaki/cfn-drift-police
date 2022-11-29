package slackcli

import "github.com/slack-go/slack"

/**
SlakClientは、Slackクライアントのインタフェースとなる構造体です。
*/
type SlackClient interface {
	/**
	PostMessageは、対象となるSlackチャンネルにメッセージを投稿する関数です。
	*/
	PostMessage(channelName string, options ...slack.MsgOption) error
}
