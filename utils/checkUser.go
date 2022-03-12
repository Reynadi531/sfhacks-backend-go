package utils

import (
	"github.com/Reynadi531/sfhacks-backend-go/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUser(email string) (bool, error) {
	_, err := database.FindUserByEmail(email)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
