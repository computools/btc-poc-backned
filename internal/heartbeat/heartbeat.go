package heartbeat

import "context"

type Service struct {
	db Database
}

func (s *Service) Ping(ctx context.Context) error {
	return s.db.Ping(ctx)
}

func NewService(db Database) *Service {
	return &Service{db: db}
}
