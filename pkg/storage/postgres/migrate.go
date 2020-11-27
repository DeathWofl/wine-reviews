package postgres

import (
	"github.com/deathwofl/wine-reviews/graph/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// drop table if exists
	db.Migrator().DropTable(
		model.Winery{},
		model.User{},
		model.Wine{},
		model.Review{},
	)

	// migrate tables and foreign key
	db.Migrator().AutoMigrate(
		model.Winery{},
		model.User{},
		model.Wine{},
		model.Review{},
	)

	return nil
}
