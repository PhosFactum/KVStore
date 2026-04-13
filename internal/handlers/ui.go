// Operations for menu-points processing
package handlers

import (
	"fmt"

	"github.com/PhosFactum/KVStore/pkg/input"
)

// ShowMenu: function to show whole menu in cli
func ShowMenu() {
	fmt.Println("            --- KVStore ---             ")
	fmt.Println("[ Program for key-value store for data ]")
	fmt.Println()
	ShowHelp()

	for {
		input, err := input.GetString()
		if err != nil {
			fmt.Println("error while writing command")
			continue
		}

		switch input {
		case "EXIT", "exit", "Exit":
			fmt.Println("\n--- Program was terminated! ---")
			return
		case "HELP", "help", "Help":
			ShowHelp()
		case "SET":
			callSET()
		case "GET":
			callGET()
		case "DELETE":
			callDELETE()
		case "STATS":
			callSTATS()
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
