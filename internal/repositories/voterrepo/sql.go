package voterrepo

import (
	"election-service/internal/core/models"

	"gorm.io/gorm"
)

type repositorySQL struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) repositorySQL {
	return repositorySQL{db: db}
}

// Retrieving objects with primary key
func (r repositorySQL) FindByID(id int) (*models.Voter, error) {
	var entity models.Voter
	tx := r.db.Find(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, nil
	}

	return &entity, nil
}

// Update with primary key
func (r repositorySQL) UpdateByID(id int, entity models.Json) (int, error) {
	tx := r.db.Model(&models.Voter{}).Where("id = ?", id).Updates(&entity)
	if tx.Error != nil {
		return int(tx.RowsAffected), tx.Error
	}

	return int(tx.RowsAffected), nil
}
