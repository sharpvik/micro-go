package service

type Users interface {
	Add(username, password string) error
}
