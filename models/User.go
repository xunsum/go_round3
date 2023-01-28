package models

type User struct {
	Id                string
	UserName          string
	Email             string
	EncryptedPassword string
	GenerateTime      int64
}

func (User) TableName() string {
	return "userdb"
}
