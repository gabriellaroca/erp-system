package repository

import "erp-system.com/v1/domain/model"

type PersonRepository interface {
	ReadByEmail(email string) (model.Person, error)
}
