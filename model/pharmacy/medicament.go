package pharmacy

import (
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"

	"farmaesro/model/entity"
)

type MedicamentSaver interface {
	// Save saves the medicament into storage.
	// TODO: посмотреть как в кибси, паттерн Repository
	// TODO: передавать по ссылке
	Save(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error)
}

type MedicamentFetcher interface {
	// Fetch fetches the medicament from storage.
	// TODO: делать как в кибси (рекомендация)
	Fetch(ctx context.Context, UUID uuid.UUID) (entity.Medicament, error)
}

type MedicamentDeleter interface {
	// Delete deletes the medicament from storage.
	// TODO: подумать о необходимости возврата объекта
	// TODO: удаление не предполагает постфактум информацию, бесполезное
	Delete(ctx context.Context, UUID uuid.UUID) (entity.Medicament, error)
}

type MedicamentUpdater interface {
	// Update updates the medicament in storage.
	// TODO: бесполезный возврат объекта
	Update(ctx context.Context, medicament entity.Medicament) (entity.Medicament, error)
}

type MedicamentRepository interface {
	MedicamentSaver
	MedicamentFetcher
	MedicamentDeleter
	MedicamentUpdater
}

// TODO: объединить сценарии и сущность
// TODO: сделать справочником, подумать над этим
// TODO: завести журнал прихода/ухода медикаментов
// TODO: доработка use-cases
// TODO: задуматься о гонках
type MedicamentCases struct {
	repository MedicamentRepository
}

func NewMedicamentCases(repository MedicamentRepository) *MedicamentCases {
	return &MedicamentCases{repository}
}

func (mc *MedicamentCases) SaveMedicament(ctx context.Context, medicament *entity.Medicament) error {
	if err := medicament.Validate(); err != nil {
		return err
	}
	_, err := mc.repository.Fetch(ctx, medicament.UUID)
	switch {
	case errors.Is(err, ErrMedicamentNotFound):
		return mc.addNewMedicament(ctx, medicament)
	case errors.Is(err, nil):
		return mc.updateExistingMedicament(ctx, medicament)
	default:
		return err
	}
}

func (mc *MedicamentCases) addNewMedicament(ctx context.Context, medicament *entity.Medicament) error {
	saveResult, err := mc.repository.Save(ctx, *medicament)
	if err != nil {
		return err
	}
	medicament.Fill(saveResult)
	return nil
}

func (mc *MedicamentCases) updateExistingMedicament(ctx context.Context, medicament *entity.Medicament) error {
	updateResult, err := mc.repository.Update(ctx, *medicament)
	if err != nil {
		return err
	}
	medicament.Fill(updateResult)
	return nil
}

func (mc *MedicamentCases) FetchMedicament(ctx context.Context, medicament *entity.Medicament) error {
	fetchResult, err := mc.repository.Fetch(ctx, medicament.UUID)
	if err != nil {
		return err
	}
	medicament.Fill(fetchResult)
	return nil
}

func (mc *MedicamentCases) DeleteMedicament(ctx context.Context, medicament *entity.Medicament) error {
	deleteResult, err := mc.repository.Delete(ctx, medicament.UUID)
	if err != nil {
		return err
	}
	medicament.Fill(deleteResult)
	return nil
}
