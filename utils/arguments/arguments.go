package arguments

import (
	"errors"
	"reporeport/utils/output"
	"slices"
)

type Argument struct {
	Argument    string
	Description string
}

var Arguments = []Argument{
	{Argument: "--help, -h", Description: "Show help information"},
	{Argument: "--version, -v", Description: "Show version information"},
	{Argument: "--include-libs", Description: "Include library files (such as node_modules or .venv) in the output"},
	{Argument: "--use-gitignore", Description: "Respect .gitignore files when generating the report"},
}

func GetArgumentDescription(argument string) string {
	if !slices.ContainsFunc(Arguments, func(arg Argument) bool { return arg.Argument == argument }) {
		output.PrintErrorAndExit(errors.New("Argument not found: " + argument))
		return ""
	}
	for _, arg := range Arguments {
		if arg.Argument == argument {
			return arg.Description
		}
	}

	// Return empty string to make the compiler happy
	// fuck compilers, honestly
	return ""
}
