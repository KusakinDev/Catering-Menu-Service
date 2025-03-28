package main

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DELETE FROM allergies;`,
		`DELETE FROM dishes;`,
		`DELETE FROM nutrition_facts;`,
		`DELETE FROM categories;`,
		`DELETE FROM allergy_types;`,
		`DELETE FROM tags;`,
		`DELETE FROM types;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing clear: ", stmt, err)
		}
	}

	log.Println("All table is clear")

	db.CloseDB()
}
