// usecase/person_usecase.go
package usecase

import (
	"erp-system.com/v1/domain/model"
	"erp-system.com/v1/domain/repository"
)

type PersonUsecase interface {
	ReadByEmail(email string) (model.Person, error)
}

type personUsecase struct {
	repo repository.PersonRepository
}

func NewPersonUsecase(repo repository.PersonRepository) PersonUsecase {
	return &personUsecase{repo: repo}
}

func (uc *personUsecase) ReadByEmail(email string) (model.Person, error) {
	return uc.repo.ReadByEmail(email)
}
