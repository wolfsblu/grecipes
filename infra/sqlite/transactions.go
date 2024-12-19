package sqlite

import "context"

func (s *Store) Begin(ctx context.Context) error {
	tx, err := s.con.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	s.tx = tx
	s.qtx = s.db.WithTx(tx)
	return nil
}

func (s *Store) Commit() error {
	defer func() {
		s.qtx = nil
	}()
	return s.tx.Commit()
}

func (s *Store) Rollback() {
	defer func() {
		s.qtx = nil
	}()
	_ = s.tx.Rollback()
}

func (s *Store) query() *Queries {
	if s.qtx != nil {
		return s.qtx
	}
	return s.db
}
