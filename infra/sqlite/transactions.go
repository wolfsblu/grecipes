package sqlite

import (
	"context"
	"github.com/wolfsblu/go-chef/domain"
)

func (s *Store) Begin(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.WrapError(err, "failed to establish transaction", domain.ErrPersistence)
	}
	s.tx = tx
	s.qtx = s.q.WithTx(tx)
	return nil
}

func (s *Store) Commit() error {
	err := s.tx.Commit()
	if err != nil {
		return domain.WrapError(err, "failed to commit transaction", domain.ErrPersistence)
	}
	return nil
}

func (s *Store) Rollback() {
	defer func() {
		s.qtx = nil
		s.tx = nil
	}()
	_ = s.tx.Rollback()
}

func (s *Store) query() *Queries {
	if s.qtx != nil {
		return s.qtx
	}
	return s.q
}
