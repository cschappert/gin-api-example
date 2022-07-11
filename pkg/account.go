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

package api

import "strings"

type Account struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	TeamID   int    `json:"team_id"`
	TeamName string `json:"team_name"`
}

// AccountService makes use of the AccountRepository interface to interact with the storage implentation.
// The actual implementation can be written to use an RDB, for example, or could also be mocked out for unit testing
// purposes.
type AccountRepository interface {
	GetAccount(id int) (*Account, error)
	ListAccounts() ([]*Account, error)
	CreateAccount(a *Account) error
	DeleteAccount(id int) error
}

type AccountService struct {
	r AccountRepository
}

// Takes an AccountRepository implementation and returns an AccountService.
func NewAccountService(r AccountRepository) AccountService {
	return AccountService{r: r}
}

// Gets an account by ID.
func (s *AccountService) GetAccount(id int) (*Account, error) {
	a, err := s.r.GetAccount(id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Lists all accounts.
func (s *AccountService) ListAccounts() ([]*Account, error) {
	accounts, err := s.r.ListAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Creates a new account.
func (s *AccountService) CreateAccount(a *Account) error {
	err := s.r.CreateAccount(a)
	return err
}

// Deletes an account by ID.
func (s *AccountService) DeleteAccount(id int) error {
	err := s.r.DeleteAccount(id)
	return err
}

// Gets an account from the repository and coverts all of its fields to uppercase before returning.
// As converting to uppercase is not infrastructure dependent, the conversion is performed in the service layer.
func (s *AccountService) GetAccountUppercase(id int) (*Account, error) {
	// Use the interface to the storage layer to get the account
	a, err := s.r.GetAccount(id)
	if err != nil {
		return nil, err
	}

	// Perform the application logic here (rather than in the storage layer, for example)
	a.Name = strings.ToUpper(a.Name)
	a.Email = strings.ToUpper(a.Email)
	a.TeamName = strings.ToUpper(a.TeamName)

	return a, nil
}
