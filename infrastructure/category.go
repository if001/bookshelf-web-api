package infrastructure

import (
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
	"bookshelf-web-api/infrastructure/tables"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{ DB : db }
}

func (r *categoryRepository) Get() (*model.Category, service.RecodeNotFoundError) {
	var category model.Category
	var err  = r.DB.Find(&category).Error
	return &category, err
}

func (r *categoryRepository) GetByBookId(bookId int64) (*[]model.Category, error) {
	var categoriesTable  []tables.Category
	var categoryModel  model.Category
	var categoryModels []model.Category

	err := r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
		Where("book_id = ?", bookId).
		Find(&categoriesTable).
		Error
	for i := range categoriesTable {
		categoryModel.Fill(
			categoriesTable[i].ID,
			categoriesTable[i].Name,
			categoriesTable[i].CreatedAt,
			categoriesTable[i].UpdatedAt,
		)
		categoryModels = append(
			categoryModels,
			categoryModel)
	}
	return &categoryModels, err
}

