package auth

import "context"

type Service struct {
	auth Auth
}

func (a *Service) Login(ctx context.Context, username, password string) (JWT, error) {
	token, err := a.auth.Login(ctx, username, password)
	if err != nil {
		return JWT{}, err
	}

	return JWT{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a *Service) RefreshToken(ctx context.Context, refreshToken string) (JWT, error) {
	token, err := a.auth.RefreshToken(ctx, refreshToken)
	if err != nil {
		return JWT{}, err
	}

	return JWT{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a *Service) Logout(ctx context.Context, refreshToken string) error {
	return a.auth.Logout(ctx, refreshToken)
}

func NewService(auth Auth) *Service {
	return &Service{auth: auth}
}
