package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func GenModels() {
	dsn := "root:my-secret-pw@tcp(localhost:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "internal/model",
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
    })

	g.UseDB(db)
	users := g.GenerateModel("users")

	room_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"room_user"},
			"references": []string{"HostID"},
		},
	}
	rooms := g.GenerateModel("rooms", gen.FieldRelate(field.HasOne, "Users", users, &room_options))

	// messages := g.GenerateModel("messages")
	// room_members := g.GenerateModel("room_members")

	g.ApplyBasic(users, rooms)

	g.Execute()
}

func main() {
	GenModels()
}
