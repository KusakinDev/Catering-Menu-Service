package main

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DROP TABLE allergies CASCADE;`,
		`DROP TABLE dishes CASCADE;`,
		`DROP TABLE nutrition_facts CASCADE;`,
		`DROP TABLE categories CASCADE;`,
		`DROP TABLE allergy_types CASCADE;`,
		`DROP TABLE tags CASCADE;`,
		`DROP TABLE types CASCADE;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing drop: ", stmt, err)
		}
	}

	log.Println("All table is droped")

	db.CloseDB()
}
