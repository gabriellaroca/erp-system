package mongodb

import (
	"context"

	"erp-system.com/v1/domain/model"
	"erp-system.com/v1/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type personRepository struct {
	collection *mongo.Collection
}

func NewPersonRepository(collection *mongo.Collection) repository.PersonRepository {
	return &personRepository{collection: collection}
}

func (repo *personRepository) ReadByEmail(email string) (model.Person, error) {
	var person model.Person
	filter := bson.M{"email": email}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Person{}, repository.ErrNotFound
		}
		return model.Person{}, err
	}

	return person, nil
}
