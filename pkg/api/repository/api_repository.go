package repository

import (
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// APIRepository repository
type APIRepository interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	Create(api domain.API) (string, error)
	Update(api domain.API) error
	Delete(id string) error
}

type apiRepository struct {
	db *gorm.DB
}

// NewAPIRepository is init for APIController
func NewAPIRepository(db *database.DB) APIRepository {
	return &apiRepository{
		db: db.Connection,
	}
}

// GetAll Get all apisdata
func (r *apiRepository) GetAll() ([]domain.API, error) {
	apis := []domain.API{}
	err := r.db.Find(&apis).Error

	return apis, err
}

// GetByID Get single apisdata
func (r *apiRepository) GetByID(id string) (domain.API, error) {
	api := domain.API{}
	err := r.db.Where("id = ?", id).First(&api).Error

	return api, err
}

// Create Add api
func (r *apiRepository) Create(api domain.API) (string, error) {
	err := r.db.Create(&api).Error
	id := api.ID
	return id, err
}

// Update Update api
func (r *apiRepository) Update(api domain.API) error {
	targetAPI := domain.API{}

	err := r.db.Where("id = ?", api.ID).First(&targetAPI).Error
	if err != nil {
		return err
	}

	return r.db.Save(&api).Error
}

// Delete Delete apidata
func (r *apiRepository) Delete(id string) error {
	api := domain.API{}

	api.ID = id
	result := r.db.Delete(&api)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
