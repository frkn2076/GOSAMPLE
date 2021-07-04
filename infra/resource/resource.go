package resource

import (
	"fmt"

	"app/GoSample/config/cache"
	"app/GoSample/db/repo"
	"app/GoSample/infra/constant"
	"app/GoSample/logger"
)

type Resource struct{
	LocalizationRepo repo.ILocalizationRepo
}

func CacheFeeder() {
	if localizations, err := repo.Localization.GetAll(); err != nil {
		errorMessage := fmt.Sprintf("An error occured while loading localizations to cache: %s", err.Error())
		logger.ErrorLog(errorMessage)
	} else {
		for _, localization := range localizations {
			key := generateResourceKey(localization.Resource, localization.Language)
			cache.Set(key, localization.Message, 0)
		}
	}
}

func (res *Resource) GetResource(resource string, language string) string {
	key := generateResourceKey(resource, language)
	message := cache.Get(key)
	if message != constant.EmptyString {
		return message
	}

	localization, err := res.LocalizationRepo.First(resource, language)
	if err != nil {
		errorMessage := fmt.Sprintf("Resource not found in database key: %s - Error: %s", key, err.Error())
		logger.ErrorLog(errorMessage)
		return constant.EmptyString
	}

	message = localization.Message
	res.SetResource(resource, language, message)
	return message
}

func (res *Resource) SetResource(resource string, language string, message string) {
	key := generateResourceKey(resource, language)
	cache.Set(key, message, 0)
}

func generateResourceKey(resource string, language string) string {
	key := fmt.Sprintf("%s-%s", resource, language)
	return key
}
