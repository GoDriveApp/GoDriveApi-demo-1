package main

import (
	ent "github.com/GoDriveApp/GoDriveApi/core/entities"
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL := "postgres://postgres:postgrespw@localhost:32768/GoDriveDb"
	db, _ := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	_ = db.AutoMigrate(&ent.Account{})
	//db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&ent.Account{})

	id := "1234567890"
	_, username := val.NewUsername("vu")
	_, passwordHash := val.NewPasswordHash("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	_, email := val.NewEmail("vupham@gmail.com")
	acc := ent.NewAccount(id, username, email, passwordHash)

	db.Create(acc)
}
