// Copyright 2021 Chris Schappert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
