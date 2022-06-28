package app

import "github.com/Inabiel/bookstore_user_api_go/controllers/users"

func MapUrls() {
	r.GET("/users", users.GetUser)
	r.GET("/user/search", users.SearchUser)
	r.GET("/user/:id", users.FindUser)
	r.POST("/user", users.CreateUser)
}
