// Copyright 2022 Chris Schappert
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
	"time"

	api "github.com/cschappert/gin-api-example/pkg"
	"gorm.io/gorm"
)

// AccountService represents a MySql implementation of api.AccountService.
type AccountService struct {
	DB *gorm.DB
}

// Maintain an Account model just for use with the DB to prevent coupling between the service layer
// and the infra layer. mysql.Account (the DB table model) can be converted to an
// api.Account (the business object) using its toEntity method.
type Account struct {
	// By default, GORM expects the primary key to be named 'id' in the table and 'ID' in the struct
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time // by default, GORM understands CreatedAt to correspond to DB column created_at
	UpdatedAt time.Time
	TeamID    int  // GORM understands TeamID to be an FK to Team.ID
	Team      Team // related data is held in this embedded struct
}

type Team struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *AccountService) GetAccount(id int) (*api.Account, error) {
	var account Account

	result := s.DB.Preload("Team").First(&account, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return account.toEntity(), nil
}

func (s *AccountService) ListAccounts() ([]*api.Account, error) {
	var accounts []Account

	result := s.DB.Preload("Team").Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	results := make([]*api.Account, 0, len(accounts))

	for _, v := range accounts {
		results = append(results, v.toEntity())
	}

	return results, nil
}

func (s *AccountService) CreateAccount(a *api.Account) error {
	account := Account{
		Name:  a.Name,
		Email: a.Email,
	}

	result := s.DB.Create(&account)

	return result.Error
}

func (s *AccountService) DeleteAccount(id int) error {
	result := s.DB.Delete(&Account{}, id)
	return result.Error
}

// Transforms a mysql.Account DB table model to an api.Account business object
func (a *Account) toEntity() *api.Account {
	account := api.Account{
		Id:       a.ID,
		Name:     a.Name,
		Email:    a.Email,
		TeamID:   a.Team.ID,
		TeamName: a.Team.Name,
	}

	return &account
}
