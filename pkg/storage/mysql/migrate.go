package mysql

import (
	"github.com/deathwofl/wine-reviews/pkg"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// drop table if exists
	db.Migrator().DropTable(
		pkg.Winery{},
		pkg.User{},
		pkg.Wine{},
		pkg.Review{},
	)

	// migrate tables and foreign key
	db.Migrator().AutoMigrate(
		pkg.Winery{},
		pkg.User{},
		pkg.Wine{},
		pkg.Review{},
	)

	return nil
}
