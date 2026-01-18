package help

import (
	"fmt"
	"os"
	"reporeport/utils/arguments"
)

func PrintHelp() {
	fmt.Printf("Usage: %s [OPTIONS] <command> [arguments]\n\n", os.Args[0])
	fmt.Println("Options:")

	for _, arg := range arguments.Arguments {
		fmt.Printf("  %-20s %s\n", arg.Argument, arg.Description)
	}
}
