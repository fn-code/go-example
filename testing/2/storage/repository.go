package storage

type UserInfo struct {
	ID     string
	Name   string
	Gender string
	Email  string
}
type UserRepo interface {
	Save(info UserInfo) error
	FindByEmail(email string) (*UserInfo, error)
}

