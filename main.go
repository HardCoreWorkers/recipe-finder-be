package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/HardCoreWorkers/recipe-finder-be/recipes"

	"github.com/gorilla/mux"
)

type Recipe struct {
	Name        string
	PrepTime    string
	CookTime    string
	ImageURL    string
	Ingredients []string
}

var recipeList []Recipe

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/recipes", returnRecipes)
	myRouter.HandleFunc("/recipe/{name}", returnRecipe)
	myRouter.HandleFunc("/recipes/update", updateRecipies)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API - Mux Routers")
	recipeList = []Recipe{
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
	json.NewEncoder(w).Encode(recipeList)
}

func returnRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnRecipe")
	vars := mux.Vars(r)
	key := vars["name"]

	for _, recipe := range recipeList {
		if recipe.Name == key {
			json.NewEncoder(w).Encode(recipeList)
		}
	}
}

func updateRecipies(w http.ResponseWriter, r *http.Request) {
	recipeList, err := recipes.GetAllRecipes()

	if err != nil {
		panic(err)
	}

	result, err := json.Marshal(recipeList)
	fmt.Fprintf(w, "Recipes Updated!")
	json.NewEncoder(w).Encode(result)
}
