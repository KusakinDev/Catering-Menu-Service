package main

import (
	"log"

	"github.com/KusakinDev/Catering-Menu-Service/internal/database"
	allergymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/allergy_model"
	allergytypemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/allergy_type_model"
	categorymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/category_model"
	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	nutritionfactmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/nutrition_fact_model"
	tagmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/tag_model"
	typemodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/type_model"
)

func seed() {
	var db database.DataBase
	db.InitDB()

	sqlStatements := []string{
		`INSERT INTO types (type_dish) VALUES
        ('Суп'),
        ('Горячее'),
        ('Закуска'),
        ('Десерт'),
        ('Салат');`,

		`INSERT INTO tags (tag_dish) VALUES
        ('Хит'),
        ('Сезонное Блюдо'),
        ('Стоп Лист'),
        ('Новинка'),
        ('Специальное Предложение');`,

		`INSERT INTO allergy_types (type) VALUES
        ('Лактоза'),
        ('Орехи'),
        ('Соя'),
        ('Рыба'),
        ('Глютен');`,

		`INSERT INTO categories (category_dish) VALUES
        ('Вегетарианское'),
        ('Веганское'),
        ('Кошерное'),
        ('Безглютеновое'),
        ('Органическое');`,

		`INSERT INTO nutrition_facts (calories, proteins, fats, carbohydrates) VALUES
        (150, 5, 2, 25),
        (350, 20, 15, 50),
        (250, 10, 10, 40),
        (400, 30, 20, 60),
        (100, 5, 1, 15);`,

		`INSERT INTO dishes (name, type_id, category_id, nut_fact_id, tag_id, recipe) VALUES
		('Томатный суп', 1, 1, 1, 2, 'Смешать помидоры, добавить специи и варить 30 минут.'),
		('Куриное филе', 2, 3, 2, 1, 'Обжарить куриные филе до золотистой корочки с добавлением специй.'),
		('Сырные палочки', 3, 2, 3, 4, 'Обвалять сыр в панировке и обжарить до румяности.'),
		('Шоколадный торт', 4, 1, 2, 3, 'Смешать все ингредиенты и выпекать в духовке 45 минут.'),
		('Греческий салат', 5, 4, 1, 5, 'Смешать свежие овощи с оливками и фетой, заправить оливковым маслом.'),
		('Сливочный суп из морепродуктов', 1, 5, 3, 3, 'Приготовить бульон с морепродуктами и сливками.'),
		('Говяжий стейк', 2, 3, 1, 2, 'Обжарить стейк до желаемой степени прожарки, добавить соль и перец.'),
		('Фруктовый салат', 5, 2, 2, 4, 'Смешать свежие фрукты и заправить йогуртом.'),
		('Картофельные зразы', 3, 1, 3, 1, 'Приготовить начинку и завернуть в картофельное тесто, обжарить.'),
		('Ягодный мусс', 4, 5, 1, 5, 'Смешать ягоды с сахаром и сливками, охладить до застывания.');`,

		`INSERT INTO allergies (dish_id, allergy_type_id) VALUES
        (1, 1),
        (2, 2),
        (3, 3),
        (4, 5),
        (5, 1),
        (6, 4),
        (7, 5),
        (8, 3),
        (9, 2),
        (10, 4);`,
	}

	for _, stmt := range sqlStatements {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing seed: ", stmt, err)
		}
	}

	log.Println("Success seeding")

	db.CloseDB()
}

func main() {

	var db database.DataBase
	db.InitDB()

	var dish dishmodel.Dish
	dish.MigrateToDB(db)

	var all allergymodel.Allergy
	all.MigrateToDB(db)

	var allType allergytypemodel.AllergyType
	allType.MigrateToDB(db)

	var cat categorymodel.Category
	cat.MigrateToDB(db)

	var nutFuct nutritionfactmodel.NutritionFact
	nutFuct.MigrateToDB(db)

	var tag tagmodel.Tag
	tag.MigrateToDB(db)

	var typ typemodel.Type
	typ.MigrateToDB(db)

	db.CloseDB()

	seed()
}
