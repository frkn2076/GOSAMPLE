package mocks

import (
	"app/GoSample/db/entities"

	"gorm.io/gorm"
)

type MockLocalizationRepo struct{}

func (e MockLocalizationRepo) Create(db *gorm.DB, localization entities.Localization) {}

func (e MockLocalizationRepo) First(resource string, language string) (entities.Localization, error) {
	return entities.Localization{ Resource: "DummyResource", Message: "DummyMessage", Language: "TR"}, nil
}

func (e MockLocalizationRepo) Update(db *gorm.DB, resource string, message string, language string) {}

func (e MockLocalizationRepo) Delete(db *gorm.DB, resource string, language string) {}

func (e MockLocalizationRepo) GetAll() ([]entities.Localization, error) {
	return []entities.Localization{
		entities.Localization{ Resource: "DummyResource1", Message: "DummyMessage1", Language: "TR" },
		entities.Localization{ Resource: "DummyResource2", Message: "DummyMessage2", Language: "TR" },
	}, nil
}
