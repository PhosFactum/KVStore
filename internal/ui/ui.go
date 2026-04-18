// Operations for menu-points processing
package ui

import (
	"fmt"
	"strings"

	h "github.com/PhosFactum/KVStore/internal/handlers"
	"github.com/PhosFactum/KVStore/pkg/input"
)

// ShowMenu: function to show whole menu in cli
func ShowMenu() {
	fmt.Println("            --- KVStore ---             ")
	fmt.Println("[ Program for key-value store for data ]")
	fmt.Println()
	ShowHelp()

	for {
		line, err := input.GetString()
		if err != nil {
			fmt.Println("error while reading command")
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0])
		args := parts[1:]

		switch cmd {
		case "EXIT":
			fmt.Println("\n--- Program was terminated! ---")
			return
		case "HELP":
			ShowHelp()
		case "SET":
			fmt.Println(h.CallSET(args))
		case "GET":
			fmt.Println(h.CallGET(args))
		case "DELETE":
			fmt.Println(h.CallDELETE(args))
		case "STATS":
			fmt.Println(h.CallSTATS(args))
		default:
			fmt.Println("Undefined method, try again!")
			fmt.Println("Type 'HELP' to see all comands")
		}

	}
}

// ShowHelp: helping message with commands
func ShowHelp() {
	fmt.Println()
	fmt.Println("      How to use KVStore:      ")
	fmt.Println("-------------------------------")
	fmt.Println("| SET: set value by key       |")
	fmt.Println("| GET: get value by key       |")
	fmt.Println("| DELETE: delete value by key |")
	fmt.Println("| STATS: show statistics      |")
	fmt.Println("| HELP: show all commands     |")
	fmt.Println("| EXIT: terminate the program |")
	fmt.Println("-------------------------------")
	fmt.Println()
}
