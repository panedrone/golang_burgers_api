package main

import (
	"app/internal/dbal"
	"app/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	err := dbal.OpenDb()
	if err != nil {
		panic(err)
	}

	err = dbal.AutoMigrate()
	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	jsonPath := fmt.Sprintf("%s/cmd/migrate/burgers.json", pwd)

	bytes, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}

	var burgers []Burger
	err = json.Unmarshal(bytes, &burgers)
	if err != nil {
		panic(err)
	}

	ingredients := map[string]int64{}

	db := dbal.Db()
	for _, burger := range burgers {
		for _, in := range burger.Ingredients {
			_, ok := ingredients[in]
			if !ok {
				m := model.Ingredient{
					Name: in,
				}
				db.Create(&m)

				ingredients[in] = m.ID
			}
		}
	}

	for _, burger := range burgers {

		var dbBurgerIngredients []model.Ingredient

		for _, in := range burger.Ingredients {

			id, ok := ingredients[in]
			if !ok {
				panic(in)
			}

			dbBurgerIngredients = append(dbBurgerIngredients, model.Ingredient{ID: id, Name: in})

		}
		dbBurger := &model.Burger{
			Name:        burger.Name,
			Description: burger.Description,
			Ingredients: dbBurgerIngredients,
		}

		dao := dbal.NewBurgersDao()

		err = dao.Create(context.Background(), dbBurger)

		//// Check for errors
		//err := db.Model(dbBurger).Association("Ingredients").Error
		//if err != nil {
		//	panic(err)
		//}
		//
		//err = db.Create(dbBurger).Error
		if err != nil {
			panic(err)
		}
	}
}

//type Burger

type Burger struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Restaurant  string   `json:"restaurant"`
	Web         string   `json:"web"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Addresses   []struct {
		AddressId int    `json:"addressId,omitempty"`
		Number    string `json:"number"`
		Line1     string `json:"line1"`
		Line2     string `json:"line2,omitempty"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		AddressID int    `json:"addressID,omitempty"`
	} `json:"addresses"`
}
