package names

import "github.com/jmoiron/sqlx"

type Repo struct {
	*sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

func (repo Repo) Add(name string) (err error) {
	_, err = repo.Exec(`INSERT INTO example (names) VALUES ($1)`, name)
	return
}
