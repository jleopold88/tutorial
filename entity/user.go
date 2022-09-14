package entity

type User struct {
	User_id int32  `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int32  `json:"age"`
}
