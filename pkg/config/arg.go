package config

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

type Args struct {
	Command string
	Entity  string
	Amount  int
	Date    string
	Id      int
}

func GetArgs() *Args {

	// args := os.Args[1:]

	// fmt.Print("Args: ", args, "\n")

	parser := argparse.NewParser("accounts", "Contractor accounting overview")

	command := parser.StringPositional(nil)
	entity := parser.StringPositional(nil)

	amount := parser.Int("a", "amount", &argparse.Options{
		Required: false,
		Default:  -1,
	})

	date := parser.String("d", "date", &argparse.Options{
		Required: false,
	})

	id := parser.Int("i", "id", &argparse.Options{
		Required: false,
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// fmt.Printf("command: %v\n", *command)
	// fmt.Printf("Amount: %v\n", *amount)
	// fmt.Printf("Entity: %v\n", *entity)
	// fmt.Printf("Date: %v\n", *date)
	// fmt.Printf("id: %v\n", *id)

	return &Args{
		Command: *command,
		Entity:  *entity,
		Amount:  *amount,
		Date:    *date,
		Id:      *id,
	}
}
