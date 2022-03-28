package domain

import (
	"github.com/nvs2394/just-bank-auth/models"
	"github.com/nvs2394/just-bank-lib/errs"
	"github.com/nvs2394/just-bank-lib/logger"
	"gorm.io/gorm"
)

type AuthRepositoryDb struct {
	client *gorm.DB
}

func (authRepo AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.AppError) {
	var login Login

	result := authRepo.client.Unscoped().Model(&models.User{}).Select("users.username as UserName, users.customer_id as CustomerId, users.role as Role, group_concat(accounts.account_id) as Accounts").Joins("LEFT JOIN accounts ON accounts.customer_id = users.customer_id").Where(models.User{UserName: username, Password: password}).Group("users.customer_id").Scan(&login)
	if result.Error != nil {
		logger.Error("Error when verifying login request from database" + result.Error.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &login, nil
}

func NewAuthRepositoryDb(dbClient *gorm.DB) AuthRepositoryDb {
	return AuthRepositoryDb{
		dbClient,
	}
}
