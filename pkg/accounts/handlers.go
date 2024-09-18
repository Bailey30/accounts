package accounts

import (
	"github.com/Bailey30/accounts/pkg/config"
	"github.com/Bailey30/accounts/pkg/db"
	"github.com/Bailey30/accounts/pkg/sale"
	"github.com/Bailey30/accounts/pkg/sale/store"
)

func EntityHandler(config *config.Config, database *db.Database) {
	switch config.Args.Entity {
	case "sale":
		sale.SaleHandler(sale.New(store.New(database.Dbx)), config)
	}
}
