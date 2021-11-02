package Commands

var (
	CommandHandler map[string]func(args *[]string) (bool, error) = map[string]func(args *[]string) (bool, error){
		"!play":   PlayCommand,
		"!stop":   StopCommand,
		"!skip":   SkipCommand,
		"!resume": ResumeCommand,
		"!pause":  PauseCommand,
	}
)
