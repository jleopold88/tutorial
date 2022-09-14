package modules

import (
	"context"
	"fmt"
	"tutorial/entity"
	"tutorial/prototype"
)

type rpc struct {
	useCase UseCase
	prototype.UnimplementedTutorialServer
}

func NewRPC(useCase UseCase) prototype.TutorialServer {
	return rpc{useCase: useCase}
}

func (r rpc) CreateUser(ctx context.Context, req *prototype.CreateData) (*prototype.SuccessMessage, error) {
	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	if err := r.useCase.CreateUser(ctx, user); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &prototype.SuccessMessage{
		Message: "User Created !",
		Code:    200,
	}, nil
}

func (r rpc) GetByID(ctx context.Context, req *prototype.GetDatabyID) (*prototype.DataReturn, error) {
	user, err := r.useCase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	proto_user := Entity2Proto(user)

	return &prototype.DataReturn{User: proto_user}, nil
}

func (r rpc) GetByName(ctx context.Context, req *prototype.GetDatabyName) (*prototype.ListDataReturn, error) {
	list_user, err := r.useCase.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	proto_list_user := Entities2Proto(list_user)
	return &prototype.ListDataReturn{Data: proto_list_user}, nil
}

func (r rpc) DeleteUser(ctx context.Context, req *prototype.DeleteData) (*prototype.SuccessMessage, error) {
	if err := r.useCase.DeleteUser(ctx, req.Id); err != nil {
		return nil, err
	}
	return &prototype.SuccessMessage{
		Message: "User deleted",
		Code:    200,
	}, nil
}

func (r rpc) UpdateUser(ctx context.Context, req *prototype.UpdateData) (*prototype.SuccessMessage, error) {

	entity_user := &entity.User{
		User_id: req.Id,
		Name:    req.Name,
		Email:   req.Email,
		Age:     req.Age,
	}

	if err := r.useCase.UpdateUser(ctx, entity_user); err != nil {
		return nil, err
	}
	return &prototype.SuccessMessage{
		Message: "User Updated",
		Code:    200,
	}, nil
}
