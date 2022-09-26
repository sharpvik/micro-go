package names

import "github.com/jmoiron/sqlx"

type Repo struct {
	*sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{db}
}

func (r Repo) Add(name string) (err error) {
	_, err = r.Exec(`INSERT INTO example (names) VALUES ($1)`, name)
	return
}
