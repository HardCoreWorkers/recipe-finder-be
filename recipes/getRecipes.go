package recipes

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math"
	"strconv"
)

type Recipe struct {
	Name        string
	PrepTime    string
	CookTime    string
	ImageURL    string
	Ingredients []string
}

var recipes []Recipe

//GetAllRecipes - Loop through a-z + 0-9
func GetAllRecipes() ([]Recipe, error) {
	for alphabet := 'a'; alphabet <= 'z'; alphabet++ {
		ScrapeRecipePages(string(alphabet))
	}
	ScrapeRecipePages("0-9")
	return recipes, nil
}

//ScrapeRecipePages - For each category, get number of pages to then loop through each
func ScrapeRecipePages(category string) {
	url := "https://www.bbc.co.uk/food/recipes/a-z/"
	fullURL := url + category
	fmt.Println(fullURL)
	doc, err := goquery.NewDocument(fullURL)

	if err != nil {
		panic(err)
	}

	pagination := doc.Find(".pagination-summary.gel-wrap b.gel-pica-bold").Text()

	// Divide number of recipes by 24 to get number of pages (24 recipes shown per page)
	numRecipes, err := strconv.ParseFloat(pagination, 10)
	numPages := math.Ceil(numRecipes / 24)
	numPg := int(numPages)

	for i := 1; i < numPg+1; i++ {
		pageNum := strconv.Itoa(i)
		fullURL := (fullURL + "/" + pageNum)
		GetRecipeURLs(fullURL)
	}
}

//GetRecipeURLs - For each recipe shown on the page, get the recipe url to then parse
func GetRecipeURLs(url string) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	doc.Find(".promo").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if ok {
			ParseRecipeDetails(link)
		}
	})
}

//ParseRecipeDetails - Parse recipe details
func ParseRecipeDetails(url string) {
	doc, err := goquery.NewDocument("https://www.bbc.co.uk" + url)

	if err != nil {
		panic(err)
	}

	name := doc.Find("h1.gel-trafalgar.content-title__text").Text()

	prepTime := doc.Find("p.recipe-metadata__prep-time").Text()

	cookTime := doc.Find("p.recipe-metadata__cook-time").Text()

	imageURL, ok := doc.Find(".recipe-media__image img").Attr("src")

	if ok {
		//yay
	}

	ingredients := []string{}
	doc.Find("ul.recipe-ingredients__list").Each(func(i int, s *goquery.Selection) {
		s.Find("a.recipe-ingredients__link").Each(func(i int, r *goquery.Selection) {
			ingredients = append(ingredients, (r.Text()))
		})
	})

	r := Recipe{
		Name:        name,
		PrepTime:    prepTime,
		CookTime:    cookTime,
		ImageURL:    imageURL,
		Ingredients: ingredients,
	}

	recipes = append(recipes, r)
}
