package colors

var ColorEscapeSequences = []string{
	"\033[31m",
	"\033[33m",
	"\033[32m",
	"\033[36m",
	"\033[34m",
	"\033[35m",
}

var ColorReset = "\033[0m"

var ColorEscapeSequencesMap = map[string]string{
	"red":     "\033[31m",
	"yellow":  "\033[33m",
	"green":   "\033[32m",
	"cyan":    "\033[36m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
}
