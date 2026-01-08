package configs

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func consoleSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		return 24, 80
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		return heigth, 80
	}
	return heigth, width
}

func GetTerminalWidth() int {
	_, width := consoleSize()
	return width
}
