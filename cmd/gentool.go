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
		ModelPkgPath: "./model",
		OutPath: 	"./internal/gorm_gen",
		OutFile:	"gentool.go",
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
    })

	g.UseDB(db)
	users := g.GenerateModel("users")

	// Define the relationship between rooms and users
	room_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"id"},
			"references": []string{"host_id"},
		},
	}
	rooms := g.GenerateModel("rooms", gen.FieldRelate(field.HasOne, "Host", users, &room_options))

	// // Define the relationship between rooms and users, room details.
	room_members_user_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"id"},
			"references": []string{"user_id"},
		},
	}
	room_members_room_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"id"},
			"references": []string{"room_id"},
		},
	}
	room_members := g.GenerateModel("room_members", 
		gen.FieldRelate(field.HasOne, "RoomMembers", users, &room_members_user_options),
		gen.FieldRelate(field.HasOne, "Room", rooms, &room_members_room_options))

	// Define the relationship between messages and users, room details.
	messages_room_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"id"},
			"references": []string{"room_id"},
		},
	}
	messages_user_options := field.RelateConfig{
		GORMTag: field.GormTag{
			"foreignKey": []string{"id"},
			"references": []string{"sender_id"},
		},
	}
	messages := g.GenerateModel("messages",
		gen.FieldRelate(field.BelongsTo, "Room", rooms, &messages_room_options),
		gen.FieldRelate(field.HasOne, "User", users, &messages_user_options))

	g.ApplyBasic(users, rooms, messages, room_members)

	g.Execute()
}

func main() {
	GenModels()
}
