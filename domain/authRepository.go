package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank-lib/errs"
	"github.com/nvs2394/just-bank-lib/logger"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (authRepo AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.AppError) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
	LEFT JOIN accounts a ON a.customer_id = u.customer_id
  WHERE username = ? and password = ?
  GROUP BY a.customer_id`

	err := authRepo.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewUnauthorizedError("Invalid credential")
		}

		logger.Error("Error when verifying login request from database" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &login, nil
}

func NewAuthRepositoryDb(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{
		dbClient,
	}
}
