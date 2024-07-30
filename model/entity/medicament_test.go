package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestMedicament_Fill(t *testing.T) {
	expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
	expectedMedicament := Medicament{
		UUID:        expectedMedicamentUUID,
		Title:       "test",
		Description: "test",
		Dosage:      100.0,
		DosageType:  DosageTypeMg,
	}
	var actualMedicament Medicament
	actualMedicament.Fill(expectedMedicament)
	assert.Equal(t, expectedMedicament, actualMedicament)
}

func TestMedicament_Validate(t *testing.T) {
	medicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))

	t.Run("valid", func(t *testing.T) {
		medicament := Medicament{
			UUID:        medicamentUUID,
			Title:       "test",
			Description: "test",
			Dosage:      100.0,
			DosageType:  DosageTypeMg,
		}
		err := medicament.Validate()
		require.NoError(t, err)
	})

	testCases := []struct {
		name       string
		medicament Medicament
	}{
		{
			"invalid title",
			Medicament{UUID: medicamentUUID, Title: "", Description: "test", Dosage: 100.0, DosageType: DosageTypeMg},
		},
		{
			"invalid description",
			Medicament{UUID: medicamentUUID, Title: "test", Description: "", Dosage: 100.0, DosageType: DosageTypeMg},
		},
		{
			"invalid dosage",
			Medicament{UUID: medicamentUUID, Title: "test", Description: "test", Dosage: 0.0, DosageType: DosageTypeMg},
		},
		{
			"invalid dosage type",
			Medicament{UUID: medicamentUUID, Title: "test", Description: "test", Dosage: 100.0, DosageType: DosageTypeUnknown},
		},
		{
			"empty invalid",
			Medicament{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.medicament.Validate()
			require.ErrorIs(t, err, ErrInvalidMedicament)
		})
	}
}
