package db

import (
	"context"
	"database/sql"
	"github.com/koko2824/first_postgresql/config"
	"github.com/koko2824/first_postgresql/models"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
)

func Init()  {
	db, err := sql.Open("postgres",config.DBConnect())
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)
}

func GetAll() models.UserSlice {
	db, err := sql.Open("postgres", config.DBConnect())
	if err != nil {
		panic(err)
	}

	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		panic(err)
	}

	return users
}

func GetOne(id int) models.User {
	db, err := sql.Open("postgres", config.DBConnect())
	if err != nil {
		panic(err)
	}

	user, err := models.FindUser(context.Background(), db, id)
	if err != nil {
		panic(err)
	}

	return *user
}

func Insert(name string, role string)  {
	db, err := sql.Open("postgres", config.DBConnect())
	if err != nil {
		panic(err)
	}

	user := models.User{
		Name: name,
		Role: role,
	}
	err = user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}
}

func Update(id int, name string, role string)  {
	db, err := sql.Open("postgres", config.DBConnect())
	if err != nil {
		panic(err)
	}

	user := models.User{
		ID: id,
		Name: name,
		Role: role,
	}
	_, err = user.Update(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}
}

func Delete(id int)  {
	db, err := sql.Open("postgres", config.DBConnect())
	if err != nil {
		panic(err)
	}

	user := models.User{ID: id}
	_, err = user.Delete(context.Background(), db)
	if err != nil {
		panic(err)
	}
}