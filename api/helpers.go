package api

import "github.com/gabielfemi/go-inform/farmerly"

func getCategories() []farmerly.Categories {
	return farmerly.FetchCategories()
}