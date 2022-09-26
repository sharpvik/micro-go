package users

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Repo struct {
	*sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

func (r Repo) Add(username, password string) (err error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	_, err = r.Exec(`
	insert into users (username, passphrase) values ($1, $2)
	`, username, hash)
	return
}
