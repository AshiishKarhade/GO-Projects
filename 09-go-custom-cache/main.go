package main

import (
	"fmt"
	"github.com/AshiishKarhade/GO-Projects/go-custom-cache/models"
)

func main() {
	fmt.Println("START CACHE")
	cache := models.NewCache()
	for _, word := range []string{"parrot", "avocado", "dragonfruit", "tree", "potato", "tomato", "tree", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
