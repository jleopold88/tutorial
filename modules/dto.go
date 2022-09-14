package modules

import (
	"tutorial/entity"
	"tutorial/prototype"
)

func Entity2Proto(user *entity.User) *prototype.User {
	return &prototype.User{
		Id:    user.User_id,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

func Entities2Proto(users []*entity.User) []*prototype.User {
	var users_proto []*prototype.User

	for _, user := range users {
		users_proto = append(users_proto, Entity2Proto(user))
	}

	return users_proto
}
