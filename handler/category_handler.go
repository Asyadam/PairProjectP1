package handler

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/entity"
	"github.com/Asyadam/PairProjectP1/repository"
)

type CategoryHandler struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryHandler(categoryRepo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{CategoryRepo: categoryRepo}
}

func (h *CategoryHandler) CreateCategory() {
	var category entity.Category

	fmt.Print("Input Category Name: ")
	fmt.Scan(&category.CategoryName)

	err := h.CategoryRepo.CreateCategory(category)

	if err != nil {
		fmt.Println("Create Category Failed:", err)
		return
	}

	fmt.Println("Create Category Success!")
}

func (h *CategoryHandler) ViewCategories() {
	categories, err := h.CategoryRepo.GetAllCategories()

	if err != nil {
		fmt.Println("Get Categories Failed:", err)
		return
	}

	fmt.Println("==== CATEGORY LIST ====")

	if len(categories) == 0 {
		fmt.Println("No categories found.")
		return
	}

	for _, category := range categories {
		fmt.Println("ID:", category.ID)
		fmt.Println("Category Name:", category.CategoryName)
		fmt.Println("=======================")
	}
}

func (h *CategoryHandler) UpdateCategory() {
	var category entity.Category

	fmt.Print("Input Category ID: ")
	fmt.Scan(&category.ID)

	fmt.Print("Input New Category Name: ")
	fmt.Scan(&category.CategoryName)

	err := h.CategoryRepo.UpdateCategory(category)

	if err != nil {
		fmt.Println("Update Category Failed:", err)
		return
	}

	fmt.Println("Update Category Success!")
}

func (h *CategoryHandler) DeleteCategory() {
	var id int

	fmt.Print("Input Category ID: ")
	fmt.Scan(&id)

	err := h.CategoryRepo.DeleteCategory(id)

	if err != nil {
		fmt.Println("Delete Category Failed:", err)
		return
	}

	fmt.Println("Delete Category Success!")
}
