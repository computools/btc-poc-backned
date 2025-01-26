package user

import (
	"btc-backend/pkg/auth"
	"btc-backend/pkg/database"
)

type CreateUser struct {
	User
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *CreateUser) toAuth() auth.User {
	return auth.User{
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	CompanyID *int64 `json:"company_id"`
	ImageURL  string `json:"image_url"`
}

func (u *User) toDatabase(keycloakID string) database.User {
	return database.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Position:   u.Position,
		CompanyID:  u.CompanyID,
		ImageURL:   u.ImageURL,
		KeycloakID: keycloakID,
	}
}

func userFromDatabase(u database.User) User {
	return User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Position:  u.Position,
		CompanyID: u.CompanyID,
		ImageURL:  u.ImageURL,
	}
}

func usersFromDatabase(u []database.User) []User {
	result := make([]User, 0, len(u))
	for _, u := range u {
		result = append(result, userFromDatabase(u))
	}

	return result
}
