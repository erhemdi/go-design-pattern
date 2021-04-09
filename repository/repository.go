package repository

import (
	"fmt"
	"log"
)

type UserRepository struct{}

func (repo *UserRepository) Insert(user User) {
	// SQL Query here.
	log.Println("Insert new user")
}

func (repo *UserRepository) Update(user User) {
	// SQL Query here.
	log.Printf("Update user %s\n", user.ID)
}

func (repo *UserRepository) Delete(id string) {
	// SQL Query here.
	log.Println("Delete user")
}

func (repo *UserRepository) FindByID(id string) User {
	return User{
		ID:      id,
		Name:    fmt.Sprintf("user %s", id),
		Balance: 100000,
	}
}
