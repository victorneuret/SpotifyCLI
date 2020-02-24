package Utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(out))
	width, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}

func GetTermHeight() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(out))
	height, err := strconv.Atoi(words[0])
	if err != nil {
		log.Fatal(err)
	}
	return height
}
