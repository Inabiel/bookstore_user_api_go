package users

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}
