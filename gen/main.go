package gen

import (
	"challenge/db"
	"challenge/models"

	"gorm.io/gen"
)

func Generate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface,
	})
	g.UseDB(db.GetConnection())

	g.ApplyBasic(models.User{})

	g.Execute()
}