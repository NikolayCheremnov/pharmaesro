package pharmacy

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"farmaesro/model/entity"
)

func TestMedicamentCases_FetchMedicament(t *testing.T) {
	ctx := context.Background()

	t.Run("medicament found", func(t *testing.T) {
		expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
		expectedMedicament := entity.Medicament{
			UUID:        expectedMedicamentUUID,
			Title:       "test",
			Description: "test",
			Dosage:      100.0,
			DosageType:  entity.DosageTypeMg,
		}
		medicamentRepository := new(medicamentRepositoryMock)
		medicamentRepository.On("Fetch", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
			UUID := args.Get(1).(uuid.UUID)
			require.Equal(t, expectedMedicamentUUID, UUID)
		}).Return(expectedMedicament, nil)

		medicamentCases := NewMedicamentCases(medicamentRepository)
		actualMedicament := entity.Medicament{
			UUID: expectedMedicamentUUID,
		}
		err := medicamentCases.FetchMedicament(ctx, &actualMedicament)
		require.NoError(t, err)
		require.Equal(t, expectedMedicament, actualMedicament)
	})

	t.Run("medicament not found", func(t *testing.T) {
		expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
		medicamentRepository := new(medicamentRepositoryMock)
		medicamentRepository.On("Fetch", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
			UUID := args.Get(1).(uuid.UUID)
			require.Equal(t, expectedMedicamentUUID, UUID)
		}).Return(entity.Medicament{}, ErrMedicamentNotFound)

		medicamentCases := NewMedicamentCases(medicamentRepository)
		actualMedicament := entity.Medicament{
			UUID: expectedMedicamentUUID,
		}
		err := medicamentCases.FetchMedicament(ctx, &actualMedicament)
		require.ErrorIs(t, err, ErrMedicamentNotFound)
	})
}

func TestMedicamentCases_SaveMedicament(t *testing.T) {
	ctx := context.Background()

	t.Run("add new medicament", func(t *testing.T) {
		expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
		expectedMedicament := entity.Medicament{
			UUID:        expectedMedicamentUUID,
			Title:       "test",
			Description: "test",
			Dosage:      100.0,
			DosageType:  entity.DosageTypeMg,
		}
		medicamentRepository := new(medicamentRepositoryMock)
		medicamentRepository.On("Fetch", mock.Anything, mock.Anything).
			Return(entity.Medicament{}, ErrMedicamentNotFound)
		medicamentRepository.On("Save", mock.Anything, mock.Anything).Return(expectedMedicament, nil)

		medicamentCases := NewMedicamentCases(medicamentRepository)
		actualMedicament := entity.Medicament{
			UUID:        uuid.Nil,
			Title:       expectedMedicament.Title,
			Description: expectedMedicament.Description,
			Dosage:      expectedMedicament.Dosage,
			DosageType:  expectedMedicament.DosageType,
		}
		err := medicamentCases.SaveMedicament(ctx, &actualMedicament)
		require.NoError(t, err)
		require.Equal(t, expectedMedicament, actualMedicament)
	})

	t.Run("update existing medicament", func(t *testing.T) {
		expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
		expectedMedicament := entity.Medicament{
			UUID:        expectedMedicamentUUID,
			Title:       "test",
			Description: "test",
			Dosage:      100.0,
			DosageType:  entity.DosageTypeMg,
		}
		medicamentRepository := new(medicamentRepositoryMock)
		medicamentRepository.On("Fetch", mock.Anything, mock.Anything).
			Return(expectedMedicament, nil)
		medicamentRepository.On("Update", mock.Anything, mock.Anything).Return(expectedMedicament, nil)

		medicamentCases := NewMedicamentCases(medicamentRepository)
		actualMedicament := entity.Medicament{
			UUID:        expectedMedicament.UUID,
			Title:       expectedMedicament.Title,
			Description: expectedMedicament.Description,
			Dosage:      expectedMedicament.Dosage,
			DosageType:  expectedMedicament.DosageType,
		}
		err := medicamentCases.SaveMedicament(ctx, &actualMedicament)
		require.NoError(t, err)
		require.Equal(t, expectedMedicament, actualMedicament)
	})
}

func TestMedicamentCases_DeleteMedicament(t *testing.T) {
	ctx := context.Background()

	t.Run("delete medicament", func(t *testing.T) {
		expectedMedicamentUUID := uuid.Must(uuid.FromString("88599ac6-13d5-44d9-8feb-879665773bf9"))
		expectedMedicament := entity.Medicament{
			UUID:        expectedMedicamentUUID,
			Title:       "test",
			Description: "test",
			Dosage:      100.0,
			DosageType:  entity.DosageTypeMg,
		}
		medicamentRepository := new(medicamentRepositoryMock)
		medicamentRepository.On("Delete", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
			UUID := args.Get(1).(uuid.UUID)
			require.Equal(t, expectedMedicamentUUID, UUID)
		}).Return(expectedMedicament, nil)

		medicamentCases := NewMedicamentCases(medicamentRepository)
		actualMedicament := entity.Medicament{
			UUID: expectedMedicamentUUID,
		}
		err := medicamentCases.DeleteMedicament(ctx, &actualMedicament)
		require.NoError(t, err)
		require.Equal(t, expectedMedicament, actualMedicament)
	})
}

type medicamentRepositoryMock struct {
	mock.Mock
}

func (m *medicamentRepositoryMock) Save(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error) {
	args := m.Called(ctx, medicament)
	return args.Get(0).(entity.Medicament), args.Error(1)
}

func (m *medicamentRepositoryMock) Fetch(ctx context.Context, UUID uuid.UUID) (entity.Medicament, error) {
	args := m.Called(ctx, UUID)
	return args.Get(0).(entity.Medicament), args.Error(1)
}

func (m *medicamentRepositoryMock) Delete(ctx context.Context, UUID uuid.UUID) (entity.Medicament, error) {
	args := m.Called(ctx, UUID)
	return args.Get(0).(entity.Medicament), args.Error(1)
}

func (m *medicamentRepositoryMock) Update(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error) {
	args := m.Called(ctx, medicament)
	return args.Get(0).(entity.Medicament), args.Error(1)
}
