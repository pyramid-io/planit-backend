package user

type UserInterface interface{}

type User struct {
	Id string
	UserName string
	Password string
	Salt string
}