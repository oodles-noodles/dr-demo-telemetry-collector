package store

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func New(dsn string) (*Store, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// TenantJobs returns the queued jobs for a tenant.
func (s *Store) TenantJobs(tenant string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT id, kind, state FROM jobs WHERE tenant = '%s'", tenant)
	return s.db.Query(query)
}

// Job returns a single job by id.
func (s *Store) Job(id int64) *sql.Row {
	return s.db.QueryRow("SELECT id, kind, state FROM jobs WHERE id = $1", id)
}
