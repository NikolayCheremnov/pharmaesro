package rdb

import (
	"context"
	"farmaesro/model/entity"
	"farmaesro/model/pharmacy"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMedicamentStorage_Fetch(t *testing.T) {
	ctx := context.Background()

	t.Run("bootstrap db test", func(t *testing.T) {
		const (
			driverName       = "postgres"
			connectionString = "postgres://postgres:postgres@localhost/pharmaestro_db?sslmode=disable"
		)
		db, err := New(driverName, connectionString)
		require.NoError(t, err)
		defer db.Close()

		medicamentStorage := NewMedicamentStorage(db)

		// Scenario:
		// 1. Add new medicament
		// 2. Fetch it
		// 3. Update
		// 4. Fetch after update
		// 5. Delete
		// 6. Fetch after delete

		t.Log("1. Add new medicament")
		medicamentUUID := uuid.Must(uuid.FromString("1c7ec1e0-ef55-4790-868d-d52bfcfdee4e"))
		newMedicament := entity.Medicament{
			UUID:        medicamentUUID,
			Title:       "парацетамол",
			Description: "жаропонижающее",
			Dosage:      100.0,
			DosageType:  entity.DosageTypeMg,
		}
		medicament, err := medicamentStorage.Save(ctx, newMedicament)
		require.NoError(t, err)
		t.Log(medicament)

		t.Log("2. Fetch it")
		medicament, err = medicamentStorage.Fetch(ctx, medicament.UUID)
		require.NoError(t, err)
		t.Log(medicament)

		t.Log("3. Update")
		medicament.Dosage = 150
		medicament, err = medicamentStorage.Update(ctx, medicament)
		require.NoError(t, err)
		t.Log(medicament)

		t.Log("4. Fetch after update")
		medicament, err = medicamentStorage.Fetch(ctx, medicament.UUID)
		require.NoError(t, err)
		t.Log(medicament)

		t.Log("5. Delete")
		medicament, err = medicamentStorage.Delete(ctx, medicament.UUID)
		require.NoError(t, err)
		t.Log(medicament)

		t.Log("6. Fetch after delete")
		_, err = medicamentStorage.Fetch(ctx, medicament.UUID)
		require.ErrorIs(t, err, pharmacy.ErrMedicamentNotFound)
		t.Log(err)
	})
}
