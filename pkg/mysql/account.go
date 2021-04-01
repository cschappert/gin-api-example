package mysql

import (
	"database/sql"

	api "github.com/cschappert/gin-api-example/pkg"
	_ "github.com/go-sql-driver/mysql"
)

// UserService represents a MySql implementation of api.AccountService.
type AccountService struct {
	DB *sql.DB
}

func (s *AccountService) GetAccount(id int) (*api.Account, error) {
	return &api.Account{}, nil
}

func (s *AccountService) ListAccounts() ([]*api.Account, error) {
	return []*api.Account{}, nil
}

func (s *AccountService) CreateAccount(a *api.Account) error {
	return nil
}

func (s *AccountService) DeleteAccount(id int) error {
	return nil
}
