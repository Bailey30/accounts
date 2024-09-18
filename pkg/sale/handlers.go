package sale

import (
	"accounts/pkg/config"
	operation "accounts/pkg/config"
	"accounts/pkg/input"
	"fmt"
	"time"
)

func SaleHandler(saleService Service, config *config.Config) error {
	switch config.Operation {
	case operation.Print:
		if config.Args.Id != -1 {
			PrintSale(saleService, config.Args.Id)
		} else {
			PrintSales(saleService)
		}
	case operation.Create:
		CreateSale(saleService)
	}

	return nil
}

func PrintSales(saleService Service) {
	sales, err := saleService.GetAll()
	if err != nil {
		fmt.Printf("Error getting all sales: %v", err.Error())
	}

	for _, sale := range sales {
		fmt.Printf("Sale ID: %d\n", sale.Id)
		fmt.Printf("Amount: $%.2f\n", sale.Amount)
		fmt.Printf("Payment Date: %s\n", sale.PaymentDate.Format("2006-01-02"))
		if sale.CreatedAt != nil {
			fmt.Printf("Created At: %s\n", sale.CreatedAt.Format(time.RFC3339))
		}
		if sale.UpdatedAt != nil {
			fmt.Printf("Updated At: %s\n", sale.UpdatedAt.Format(time.RFC3339))
		}
		fmt.Println("----------------------")
	}

}

func PrintSale(saleService Service, id int) {

}

func CreateSale(saleService Service) {
	// get the details from the input in the terminal
	amount, date, err := input.GetSaleDetails()

	if err != nil {
		fmt.Printf("Error getting sale details: %v\n", err)
	}

	fmt.Printf("Amount: %.2f, Date: %v\n", amount, date)

	// create an instance of sale
	sale := &Sale{
		Amount:      amount,
		PaymentDate: date,
	}

	// save the sale to the database
	id, err := saleService.Create(*sale)
	if err != nil {
		fmt.Printf("Error saving sale to database: %v\n", err)
	}

	// log the successfully saved sale
	fmt.Printf("Sale created with Id: %v", id)
}
