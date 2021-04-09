package facade

import (
	"log"

	"github.com/mehre/go-design-pattern/repository"
)

type BankFacade struct{}

var userRepository repository.UserRepository

func (facade BankFacade) Transfer(from string, to string, amount int) {

	fromUser := userRepository.FindByID(from)
	toUser := userRepository.FindByID(to)

	fromUser.Balance -= amount
	toUser.Balance += amount

	userRepository.Update(fromUser)
	userRepository.Update(toUser)

	log.Printf("Success Transfer from user id %s to %s\n", from, to)
}
