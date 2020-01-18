package services

import (
	"github.com/tatamiya/golang-microservices/mvc/domain"
	"github.com/tatamiya/golang-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}