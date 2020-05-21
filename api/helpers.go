package api

import "github.com/gabielfemi/go-inform/farmerly"

func getCategories() []farmerly.Categories {
	return farmerly.FetchCategories()
}
func getUsers() []farmerly.Users {
	return farmerly.FetchUsers()
}

func getPosts() []farmerly.Posts {
	return farmerly.GetPosts()
}