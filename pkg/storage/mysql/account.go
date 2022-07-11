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

package storage

import (
	"time"

	api "github.com/cschappert/gin-api-example/pkg"
	"gorm.io/gorm"
)

// storage.AccountRepository represents a MySQL implementation of api.AccountRepository.
type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{db: db}
}

// Maintain an Account model just for use with the DB to prevent coupling between the service layer
// and the storage layer. storage.Account (the DB table model) can be converted to an
// api.Account (the business object) using its toEntity method.
type Account struct {
	// Some notes on GORM conventions:
	// 1. struct name Account corresponds to table name accounts. PetOwner corresponds to pet_owners
	// 2. struct field ID corresponds to table column id and is understood to be the primary key
	// 3. struct fields are CamelCase and table columns are snake_case. CreatedAt == created_at
	// 4. TeamID is understood to refer to struct field Team.ID, or table column teams.id and is a foreign key.
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	TeamID    *int

	// related data is held in this embedded struct
	Team Team
}

type Team struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *AccountRepository) GetAccount(id int) (*api.Account, error) {
	var account Account

	result := s.db.Preload("Team").First(&account, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return account.toEntity(), nil
}

func (s *AccountRepository) ListAccounts() ([]*api.Account, error) {
	var accounts []Account

	result := s.db.Preload("Team").Find(&accounts)

	if result.Error != nil {
		return nil, result.Error
	}

	results := make([]*api.Account, 0, len(accounts))

	for _, v := range accounts {
		results = append(results, v.toEntity())
	}

	return results, nil
}

func (s *AccountRepository) CreateAccount(a *api.Account) error {

	var teamId *int = nil

	if a.TeamID > 0 {
		teamId = &a.TeamID
	}

	account := Account{
		Name:   a.Name,
		Email:  a.Email,
		TeamID: teamId,
	}

	result := s.db.Create(&account)

	return result.Error
}

func (s *AccountRepository) DeleteAccount(id int) error {
	result := s.db.Delete(&Account{}, id)
	return result.Error
}

// Transforms a storage.Account DB table model to an api.Account business object
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
