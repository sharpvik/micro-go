package service

type NamesRepo interface {
	Add(name string) error
}
