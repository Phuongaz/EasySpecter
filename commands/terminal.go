package commands

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadConsole(log *log.Logger) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if len(text) == 0 {
			continue
		}
		parameter := strings.Split(text, " ")
		if GetCommandManager().CheckCommand(parameter[0]) {
			GetCommandManager().ExecuteCommand(parameter[0], parameter[1:], log)
		} else {
			log.Println("Command not found")
		}
	}
}
