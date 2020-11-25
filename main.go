package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Recipe struct {
	Name        string
	PrepTime    string
	CookTime    string
	ImageURL    string
	Ingredients []string
}

var Recipes []Recipe

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/recipes", returnRecipes)
	myRouter.HandleFunc("/recipe/{name}", returnRecipe)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API - Mux Routers")
	Recipes = []Recipe{
		Recipe{Name: "NandosChicken", PrepTime: "2mins", CookTime: "40mins", ImageURL: "N/A", Ingredients: []string{"Chicken", "Secret Spices"}},
		Recipe{Name: "KFCChicken", PrepTime: "5mins", CookTime: "15mins", ImageURL: "N/A", Ingredients: []string{"Chicken", "Secret Spices", "Oil"}},
	}
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Recipe Finder!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnRecipes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnRecipes")
	json.NewEncoder(w).Encode(Recipes)
}

func returnRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnRecipe")
	vars := mux.Vars(r)
	key := vars["name"]

	for _, recipe := range Recipes {
		if recipe.Name == key {
			json.NewEncoder(w).Encode(recipe)
		}
	}
}
