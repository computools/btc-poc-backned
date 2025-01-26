package user

import (
	"context"
)

type Service struct {
	db   Database
	auth Auth
}

func (s *Service) GetUser(ctx context.Context, id int64) (User, error) {
	user, err := s.db.GetUser(ctx, id)
	return userFromDatabase(user), err
}

func (s *Service) GetUsersByCompanyID(ctx context.Context, id int64) ([]User, error) {
	users, err := s.db.GetUsersByCompanyID(ctx, id)
	return usersFromDatabase(users), err
}

func (s *Service) CreateUser(ctx context.Context, user CreateUser) (User, error) {
	keycloakID, err := s.auth.CreateUser(ctx, user.toAuth(), user.Password)
	if err != nil {
		return User{}, err
	}

	u, err := s.db.CreateUser(ctx, user.toDatabase(keycloakID))
	return userFromDatabase(u), err
}

func (s *Service) UpdateUser(ctx context.Context, user User) (User, error) {
	u, err := s.db.UpdateUser(ctx, user.toDatabase(""))
	return userFromDatabase(u), err
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.db.DeleteUser(ctx, id)
}

func NewService(db Database, auth Auth) *Service {
	return &Service{db: db, auth: auth}
}
