package main

import (
	"log"
	"os"

	"github.com/phuongaz/easyspecter/commands"
	"github.com/phuongaz/easyspecter/specter"
)

func main() {
	log := log.New(os.Stdout, "[SPECTER] ", log.LstdFlags)
	log.Println("-------EASY SPECTER - AUTHOR PHUONGAZ--------")
	log.Println("1. XBOX ACCOUNT LOG: Xbox account required")
	log.Println("	+ Login command: xbox login")
	log.Println("	+ Spawn command: xbox join <ip:port>")
	log.Println("2. JOIN GAME: No account required")
	log.Println("	+ Must disable xbox-auth in server")
	log.Println("	+ Spawn command: join <ip:port>")
	log.Println("3. LIST: List all specters")
	log.Println("	+ Command: list")
	log.Println("4. QUIT: Quit specter")
	log.Println("	+ Command: quit <specter name>")
	log.Println("5. SPECTER: Specter control")
	log.Println("	+ Chat Command: specter chat <specter name> <message>")
	log.Println("	+ Move Command: specter move <specter name> <x> <y> <z>")
	log.Println("---------------------------------")
	commands.InitializeConsole(log)
}
