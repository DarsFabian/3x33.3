package utility

import (
	voice "redbluecircle/voip"
)

type BotUtil struct {
	Token  string
	Queues map[string]voice.Queue
}

var BotUtils = BotUtil{Token: "", Queues: map[string]voice.Queue{}}
