package models

type User struct {
	Email    string
	PassHash []byte
	ID       int64
}
