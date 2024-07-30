package pharmacy

import "errors"

var (
	ErrMedicamentNotFound        = errors.New("medicament not found")
	ErrMedicamentDuplicateUUID   = errors.New("duplicated medicament uuid")
	ErrMedicamentDuplicatedTitle = errors.New("duplicated medicament title")
)
