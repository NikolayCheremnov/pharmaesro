package entity

import (
	"errors"
	"fmt"
	"slices"

	uuid "github.com/satori/go.uuid"
)

type DosageType int

//go:generate enumer -trimprefix=DosageType -type=DosageType -json -output dosage_type_enum.go
const (
	DosageTypeUnknown DosageType = iota
	DosageTypeMg
)

type Medicament struct {
	UUID        uuid.UUID
	Title       string
	Description string
	// TODO: дозировка есть отлельный объект (активное вещество) - отдельная стркутуру
	Dosage     float64
	DosageType DosageType
}

func (m *Medicament) Fill(other Medicament) {
	m.UUID = other.UUID
	m.Title = other.Title
	m.Description = other.Description
	m.Dosage = other.Dosage
	m.DosageType = other.DosageType
}

func (m *Medicament) Validate() error {
	validationErr := ErrInvalidMedicament
	if m.Title == "" {
		validationErr = fmt.Errorf("%w: empty title", validationErr)
	}
	if m.Description == "" {
		validationErr = fmt.Errorf("%w: empty description", validationErr)
	}
	if m.Dosage <= 0 {
		validationErr = fmt.Errorf("%w: non positive dosage", validationErr)
	}
	if m.DosageType == DosageTypeUnknown || !slices.Contains(DosageTypeValues(), m.DosageType) {
		validationErr = fmt.Errorf("%w: invalid dosage type", validationErr)
	}
	if err := errors.Unwrap(validationErr); err != nil {
		return err
	}
	return nil
}
