package db

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

func ConnectToDb() (*pg.DB, error) {
	options, err := pg.ParseURL(config.Config.DBURL)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(options)
	return db, createSchema(db)
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{} {(*models.Company)(nil), (*models.Machine)(nil), (*models.MachineInspection)(nil), (*models.User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}