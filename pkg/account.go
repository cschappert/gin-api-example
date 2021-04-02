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

package api

type Account struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Team struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contact_name"`
}

type AccountService interface {
	GetAccount(id int) (*Account, error)
	ListAccounts() ([]*Account, error)
	CreateAccount(a *Account) error
	DeleteAccount(id int) error
}

type TeamService interface {
	GetTeam(id int) (*Team, error)
	ListTeams() ([]*Team, error)
	CreateTeam(t *Team) error
	DeleteTeam(id int) error
}
