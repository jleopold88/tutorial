package modules

import (
	"context"
	"database/sql"
	"tutorial/entity"

	"github.com/pkg/errors"
)

type db struct {
	conn *sql.DB
}

type Repository interface {
	GetByID(ctx context.Context, id int32) (*entity.User, error)
	GetByName(ctx context.Context, name string) ([]*entity.User, error)
	CreateUser(ctx context.Context, name, email string, age int32) error
	UpdateUser(ctx context.Context, name, email string, age, id int32) error
	DeleteUser(ctx context.Context, id int32) error
}

func NewRepository(ds *sql.DB) Repository {
	return db{conn: ds}
}

func (c db) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	sqlQuery := `select * from "user" u where u.user_id = $1 `
	user := &entity.User{}
	err := c.conn.QueryRow(sqlQuery, id).Scan(user.User_id, user.Name, user.Email, user.Age)
	if err != nil {
		errors.Wrap(err, "Failed getting User")
		return nil, err
	}
	return user, nil
}

func (c db) GetByName(ctx context.Context, name string) ([]*entity.User, error) {
	sqlQuery := `select * from "user" u where u."name" like '%$1%'`
	res, err := c.conn.Query(sqlQuery, name)
	if err != nil {
		errors.Wrap(err, "Failed getting User")
		return nil, err
	}
	users := []*entity.User{}
	for res.Next() {
		temp := entity.User{}
		if err = res.Scan(temp.User_id, temp.Name, temp.Email, temp.Age); err != nil {
			errors.Wrap(err, "Failed getting User")
			return nil, err
		}
		users = append(users, &temp)
	}
	return users, nil
}

func (c db) CreateUser(ctx context.Context, name, email string, age int32) error {
	sqlQuery := `insert into "user" ("name" , "email", "age") values($1,$2,$3);`
	if _, err := c.conn.Exec(sqlQuery, name, email, age); err != nil {
		errors.Wrap(err, "Failed creating User")
		return err
	}
	return nil
}

func (c db) UpdateUser(ctx context.Context, name, email string, age, id int32) error {
	sqlQuery := `update "user" set name = $1, email = $2, age = $3 where user_id = $4; `
	if _, err := c.conn.Exec(sqlQuery, name, email, age, id); err != nil {
		errors.Wrap(err, "Failed updating User")
		return err
	}
	return nil
}

func (c db) DeleteUser(ctx context.Context, id int32) error {
	sqlQuery := `delete from "user" where user_id = $1; `
	if _, err := c.conn.Exec(sqlQuery, id); err != nil {
		errors.Wrap(err, "Failed deleting User")
		return err
	}
	return nil
}
