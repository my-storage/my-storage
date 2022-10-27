package entities

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (a *Account) SetName(name string) {
	a.Name = name
}

func (a *Account) SetEmail(email string) {
	a.Email = email
}

func (a *Account) SetPassword(password string) {
	a.Password = password
}
