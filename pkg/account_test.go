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

package api_test

import (
	"testing"

	"github.com/cschappert/gin-api-example/mocks"
	api "github.com/cschappert/gin-api-example/pkg"
	"github.com/google/go-cmp/cmp"
)

func TestGetAccountUppercase(t *testing.T) {

	// arg is the argument for AccountService.GetAccountUppercase.
	// mockReturns is what we would like our mock repository to return to the service layer for a given test.
	// want is what we expect AccountService to return for a given argument.
	tests := []struct {
		arg         int
		mockReturns *api.Account
		want        *api.Account
	}{
		{
			arg: 1,
			mockReturns: &api.Account{
				Id:       1,
				Name:     "Alice",
				Email:    "alice@example.com",
				TeamID:   1,
				TeamName: "team 1",
			},
			want: &api.Account{
				Id:       1,
				Name:     "ALICE",
				Email:    "ALICE@EXAMPLE.COM",
				TeamID:   1,
				TeamName: "TEAM 1",
			},
		},
		{
			arg: 2,
			mockReturns: &api.Account{
				Id:       2,
				Name:     "Bob",
				Email:    "bob@example.com",
				TeamID:   2,
				TeamName: "team 2",
			},
			want: &api.Account{
				Id:       2,
				Name:     "BOB",
				Email:    "BOB@EXAMPLE.COM",
				TeamID:   2,
				TeamName: "TEAM 2",
			},
		},
	}

	// Instantiate a mock AccountRepository
	ar := &mocks.AccountRepository{}
	// Create a new AccountService using the mock repo
	as := api.NewAccountService(ar)

	for _, test := range tests {
		ar.On("GetAccount", test.arg).Return(test.mockReturns, nil)
		got, err := as.GetAccountUppercase(test.arg)
		if err != nil {
			t.Errorf("AccountService.GetAccountUppercase returned unexpected error.")
		}
		if !cmp.Equal(got, test.want) {
			t.Errorf("AccountService.GetAccountUppercase(id). Got %v, wanted %v", got, test.want)
		}
	}
}
