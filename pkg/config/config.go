package config

import (
	"fmt"
)

type Config struct {
	Operation Operation
	Args      Args
}

// i need to be able to:
// print an overview of everything eg total amount from sales; total left after salary, dividends and expenses; total corporation tax, dividend tax and student loan so far
// print each entity eg sales, invoices, dividends, taxes, salary with a month breakdown and total
// print a selected month for each
// print a range of months for each
// delete any entity by id
// edit any entity by id
// add any sales, invoices, dividends, expenses, salary with amount and date added

var operations = map[string]Operation{
	"all":    Print,
	"add":    Create,
	"delete": Delete,
	"edit":   Edit,
}

func getOperation(args *Args) (Operation, error) {
	if operation, ok := operations[args.Command]; ok {
		return operation, nil
	}

	return InvalidOperation, fmt.Errorf("Command '%s' not recognised.", args.Command)
}

func NewConfig(args *Args) (*Config, error) {
	operation, err := getOperation(args)
	if err != nil {
		return nil, err
	}

	return &Config{
		Operation: operation,
		Args:      *args,
	}, nil
}
