package model

type UserCredential struct {
	Id           uint   `db:"id"`
	Username     string `db:"user_name"`
	IsBlocked    bool   `db:"is_blocked"`
	UserPassword string `db:"user_password"`
}
