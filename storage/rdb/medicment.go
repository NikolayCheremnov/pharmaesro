package rdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	"farmaesro/model/entity"
	"farmaesro/model/pharmacy"
)

type MedicamentStorage struct {
	client *DB
}

func NewMedicamentStorage(db *DB) *MedicamentStorage {
	return &MedicamentStorage{client: db}
}

var _ pharmacy.MedicamentRepository = (*MedicamentStorage)(nil)

func (m *MedicamentStorage) Save(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error) {
	builder := sq.Insert("medicament").
		Columns("uuid", "title", "description", "dosage", "dosage_type").
		Values(medicament.UUID, medicament.Title, medicament.Description, medicament.Dosage, medicament.DosageType)
	query, args := builder.PlaceholderFormat(sq.Dollar).MustSql()
	if err := m.runExecContextQuery(ctx, query, args...); err != nil {
		var pqErr *pq.Error
		// TODO: use named constants for errors and inline it into ddl
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			switch pqErr.Constraint {
			case "medicament_pkey":
				return entity.Medicament{}, pharmacy.ErrMedicamentDuplicateUUID
			case "medicament_title_key":
				return entity.Medicament{}, pharmacy.ErrMedicamentDuplicatedTitle
			}
		}
		return entity.Medicament{}, err
	}
	return medicament, nil
}

func (m *MedicamentStorage) Fetch(ctx context.Context, uuid uuid.UUID) (entity.Medicament, error) {
	builder := sq.Select("uuid", "title", "description", "dosage", "dosage_type").
		From("medicament").
		Where(sq.Eq{"uuid": uuid})
	query, args := builder.PlaceholderFormat(sq.Dollar).MustSql()
	var medicament entity.Medicament
	err := m.client.Session.QueryRowContext(ctx, query, args...).Scan(
		&medicament.UUID,
		&medicament.Title,
		&medicament.Description,
		&medicament.Dosage,
		&medicament.DosageType,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Medicament{}, pharmacy.ErrMedicamentNotFound
	}
	if err != nil {
		return entity.Medicament{}, fmt.Errorf("unable to execute query: %w", err)
	}
	return medicament, nil
}

func (m *MedicamentStorage) Delete(ctx context.Context, uuid uuid.UUID) (entity.Medicament, error) {
	deleted, er := m.Fetch(ctx, uuid)
	if er != nil {
		return entity.Medicament{}, er
	}
	builder := sq.Delete("medicament").Where(sq.Eq{"uuid": uuid})
	query, args := builder.PlaceholderFormat(sq.Dollar).MustSql()
	if err := m.runExecContextQuery(ctx, query, args...); err != nil {
		return entity.Medicament{}, err
	}
	return deleted, nil
}

func (m *MedicamentStorage) Update(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error) {
	_, er := m.Fetch(ctx, medicament.UUID)
	if er != nil {
		return entity.Medicament{}, er
	}
	builder := sq.Update("medicament").Where(sq.Eq{"uuid": medicament.UUID}).
		Set("title", medicament.Title).
		Set("description", medicament.Description).
		Set("dosage", medicament.Dosage).
		Set("dosage_type", medicament.DosageType)
	query, args := builder.PlaceholderFormat(sq.Dollar).MustSql()
	if err := m.runExecContextQuery(ctx, query, args...); err != nil {
		return entity.Medicament{}, err
	}
	return medicament, nil
}

func (m *MedicamentStorage) runExecContextQuery(ctx context.Context, query string, args ...interface{}) error {
	result, err := m.client.Session.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("unable to execute query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to fetch affected rows: %w", err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("%d rows affected, expected 1", rowsAffected)
	}
	return nil
}
