package auth

import (
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"

	"btc-backend/config"
)

type Keycloak struct {
	cfg    *config.KeyCloakConfig
	client *gocloak.GoCloak
}

func (k *Keycloak) CreateUser(ctx context.Context, user User, password string) (string, error) {
	adminToken, err := k.client.LoginClient(ctx, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm)
	if err != nil {
		return "", errors.Join(ErrLoginClient, err)
	}

	userID, err := k.client.CreateUser(ctx, adminToken.AccessToken, k.cfg.Realm, user.toGoCloak())
	if err != nil {
		return "", errors.Join(ErrCreateUserFailed, err)
	}

	return userID, k.client.SetPassword(ctx, adminToken.AccessToken, userID, k.cfg.Realm, password, false)
}

func (k *Keycloak) DeleteUser(ctx context.Context, userID string) error {
	adminToken, err := k.client.LoginClient(ctx, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm)
	if err != nil {
		return errors.Join(ErrLoginClient, err)
	}

	return k.client.DeleteUser(ctx, adminToken.AccessToken, k.cfg.Realm, userID)
}

func (k *Keycloak) Login(ctx context.Context, user, password string) (JWT, error) {
	tokens, err := k.client.Login(ctx, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm, user, password)
	if err != nil {
		return JWT{}, err
	}

	return fromGoCloakJWT(tokens), nil
}

func (k *Keycloak) ValidateToken(ctx context.Context, accessToken string) error {
	result, err := k.client.RetrospectToken(ctx, accessToken, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm)
	if err != nil {
		return err
	}

	if result == nil || !*result.Active {
		return ErrTokenNotActive
	}

	return nil
}

func (k *Keycloak) RefreshToken(ctx context.Context, refreshToken string) (JWT, error) {
	tokens, err := k.client.RefreshToken(ctx, refreshToken, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm)
	if err != nil {
		return JWT{}, err
	}

	return fromGoCloakJWT(tokens), nil
}

func (k *Keycloak) Logout(ctx context.Context, refreshToken string) error {
	return k.client.Logout(ctx, k.cfg.ClientID, k.cfg.ClientCredentials, k.cfg.Realm, refreshToken)
}

func (k *Keycloak) Start(_ context.Context) error {
	k.client = gocloak.NewClient(k.cfg.Addr)
	return nil
}

func (k *Keycloak) Stop(_ context.Context) error {
	return nil
}

func NewKeycloak(cfg *config.Config) *Keycloak {
	return &Keycloak{cfg: &cfg.KeyCloakConfig}
}
