package main

import (
	"log"
	"os"

	"github.com/phuongaz/easyspecter/commands"
)

func main() {
	log := log.New(os.Stdout, "[SPECTER] ", log.LstdFlags)
	log.Println("-------EASY SPECTER - AUTHOR PHUONGAZ--------")
	log.Println("1. XBOX ACCOUNT LOG: Xbox account required")
	log.Println("	+ Login command: xbox login")
	log.Println("	+ Spawn command: xbox join <ip:port>")
	log.Println("2. JOIN GAME: No account required")
	log.Println("	+ Must disable xbox-auth in server")
	log.Println("	+ Spawn command: join <name> <ip:port>")
	log.Println("5. SPECTER: Specter commands")
	log.Println("	+ Chat Command: specter chat <specter name> <message>")
	log.Println("	+ Move Command: specter move <specter name> <x> <y> <z>")
	log.Println("	+ Quit Command: specter quit <specter name>")
	log.Println("	+ List Command: specter list")
	log.Println("---------------------------------")
	commands.RegisterCommands(log)
	commands.ReadConsole(log)
}
