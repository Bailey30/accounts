package accounts

import (
	"accounts/pkg/config"
	"accounts/pkg/db"
	"accounts/pkg/sale"
	"accounts/pkg/sale/store"
)

func EntityHandler(config *config.Config, database *db.Database) {
	switch config.Args.Entity {
	case "sale":
		sale.SaleHandler(sale.New(store.New(database.Dbx)), config)
	}
}
