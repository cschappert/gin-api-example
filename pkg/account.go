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
