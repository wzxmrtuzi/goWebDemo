package entity

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Password string `json:"password"`
}