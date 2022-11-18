package service

import (
	"yy-data-processing/common/database"
	"yy-data-processing/model"
)

type UserService struct {
}

func (s *UserService) SaveUser(user model.User) error {
	dbResp := database.Db.Create(&user)
	if dbResp.Error != nil {
		return dbResp.Error
	}
	return nil
}

func (s *UserService) FindUser(id int) model.User {
	user := model.User{}
	database.Db.Where("id = ?", id).Find(&user)
	return user
}
