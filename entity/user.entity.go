package entity

import (
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userTableName = "main.users"
)

type User struct {
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Auditable
}

func (model *User) TableName() string {
	return userTableName
}

func NewUser(
	uuid uuid.UUID,
	name string,
	email string,
	password string,
	createdBy string,
) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		UUID:      uuid,
		Name:      name,
		Email:     email,
		Password:  string(passwordHash),
		Auditable: NewAuditable(createdBy),
	}
}

func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"email":      model.Email,
			"updated_at": model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Email != from.Email {
		mapped["email"] = from.Email
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
