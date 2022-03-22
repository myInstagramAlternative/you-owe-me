package repository

import (
	"casbin-golang/model"
	logger "log"

	"gorm.io/gorm"
)

type logRepository struct {
	DB *gorm.DB
}

// UserRepository : represent the user's repository contract
type LogRepository interface {
	AddLog(model.Log) (model.Log, error)
	// GetLog(int) (model.Log, error)
	// GetByHost(string) (model.Log, error)
	// GetAllLogs() ([]model.Log, error)
	// UpdateLog(model.User) (model.Log, error)
	// DeleteLog(model.User) (model.Log, error)
	Migrate() error
}

// NewUserRepository -> returns new user repository
func NewLogRepository(db *gorm.DB) LogRepository {
	return logRepository{
		DB: db,
	}
}

func (u logRepository) Migrate() error {
	logger.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&model.Log{})
}

//
// func (u logRepository) GetLog(id int) (user model.User, err error) {
// 	return user, u.DB.First(&user, id).Error
// }
//
// func (u logRepository) GetByHost(email string) (user model.User, err error) {
// 	return user, u.DB.First(&user, "email=?", email).Error
// }
//
// func (u logRepository) GetAllUser() (users []model.User, err error) {
// 	return users, u.DB.Find(&users).Error
// }
//
func (u logRepository) AddLog(log model.Log) (model.Log, error) {
	return log, u.DB.Create(&log).Error
}

//
// func (u logRepository) UpdateUser(user model.User) (model.User, error) {
// 	if err := u.DB.First(&user, user.ID).Error; err != nil {
// 		return user, err
// 	}
// 	return user, u.DB.Model(&user).Updates(&user).Error
// }
//
// func (u logRepository) DeleteUser(user model.User) (model.User, error) {
// 	if err := u.DB.First(&user, user.ID).Error; err != nil {
// 		return user, err
// 	}
// 	return user, u.DB.Delete(&user).Error
// }
