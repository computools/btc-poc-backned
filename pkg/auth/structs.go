package auth

import "github.com/Nerzal/gocloak/v13"

type User struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
}

func (u User) toGoCloak() gocloak.User {
	return gocloak.User{
		Username:      gocloak.StringP(u.Username),
		FirstName:     gocloak.StringP(u.FirstName),
		LastName:      gocloak.StringP(u.LastName),
		Email:         gocloak.StringP(u.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}
}

type JWT struct {
	AccessToken  string
	RefreshToken string
}

func fromGoCloakJWT(jwt *gocloak.JWT) JWT {
	return JWT{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
	}
}
