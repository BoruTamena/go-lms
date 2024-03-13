package repository

import (
	"database/sql"
	"log"

	"github.com/BoruTamena/habitant_hub/internal/app/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {

	return &UserRepository{
		db: db,
	}

}

func (r *UserRepository) CreateNewUser(user *model.User) error {

	query := `INSERT INTO users (username,email,password)VALUES($1,$2,$3)`

	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		log.Printf("Error in cretaing user ,%v", err)

		return err
	}

	// log.Printf(res)

	return nil

}

// func (r * UserRepository)GetUserById(id int ) (* model.User,error)
// {

// }
