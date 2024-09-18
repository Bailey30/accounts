package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetSaleDetails() (float64, time.Time, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("invalid amount: %v", err)
	}

	fmt.Print("Payment Date (YYYY-MM-DD): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)

	layout := "2006-01-02" // Layout for parsing dates in YYYY-MM-DD format
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("invalid date: %v", err)
	}

	return amount, date, nil
}
