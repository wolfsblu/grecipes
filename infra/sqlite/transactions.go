package sqlite

import (
	"context"
	"github.com/wolfsblu/go-chef/domain"
	"log"
)

func (s *Store) Begin(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("failed to establish transaction:", err)
		return &domain.ErrPersistence
	}
	s.tx = tx
	s.qtx = s.q.WithTx(tx)
	return nil
}

func (s *Store) Commit() error {
	err := s.tx.Commit()
	if err != nil {
		log.Println("failed to commit transaction:", err)
		return &domain.ErrPersistence
	}
	return nil
}

func (s *Store) Rollback() {
	defer func() {
		s.qtx = nil
		s.tx = nil
	}()
	err := s.tx.Rollback()
	if err != nil {
		log.Println("failed to rollback transaction:", err)
	}
}

func (s *Store) query() *Queries {
	if s.qtx != nil {
		return s.qtx
	}
	return s.q
}
