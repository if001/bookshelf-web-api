package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"bookshelf-web-api/infrastructure/tables"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{ DB : db }
}

func (r *categoryRepository) GetCategories() (*[]model.Category, error) {
	var category []model.Category
	 err  := r.DB.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetByIds(categoriesId []int64) (*[]model.Category, error) {
	var categories []model.Category
	err := r.DB.Find(&categories, "id IN (?)", categoriesId).Error
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, service.NewRecodeNotFoundError()
	}
	return &categories, nil
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

func (r *categoryRepository) GetNotExistCategories(categories []model.Category) (*[]model.Category, error) {
	notExistCategories := []model.Category{}
	bindCategory := []model.Category{}
	for i := range categories {
		err := r.DB.Where("name = ?",categories[i].Name).Find(&bindCategory).Error
		if err != nil {
			return nil, err
		}
		if len(bindCategory) == 0 {
			notExistCategory := model.Category{}
			notExistCategory.Name = categories[i].Name
			notExistCategories = append(
				notExistCategories,
				notExistCategory,
			)
		}
	}
	return &notExistCategories, nil
}

func  (r *categoryRepository) LogicalDelete(bookId int64, categoryId int64)  (error){
	bookCategory := []tables.BookCategory{}
	err := r.DB.
		Where("book_id = ?", bookId).
		Where("category_id = ?", categoryId).
		Find(&bookCategory).Error

	if err != nil {
		return err
	}
	if len(bookCategory) == 0 {
		return service.NewRecodeNotFoundError()
	}
	bookCategory[0].Status = pFalse

	err = r.DB.Save(&bookCategory[0]).Error
	if err != nil {
		return err
	}
	return nil
}