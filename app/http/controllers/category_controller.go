package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type CategoryController struct {
	//Dependent services
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		//Inject services
	}
}

func (r *CategoryController) Index(ctx http.Context) http.Response {
	var categories []models.Category

	if err := facades.Orm().Query().Find(&categories); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Data fetch successfully",
		"data":    categories,
	})
}

func (r *CategoryController) Show(ctx http.Context) http.Response {
	category := models.Category{}

	if err := facades.Orm().Query().Where("id", ctx.Request().Input("id")).First(&category); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Category fetch successfully",
		"data":    category,
	})
}

func (r *CategoryController) Store(ctx http.Context) http.Response {
	if err := facades.Orm().Query().Create(&models.Category{
		Name: ctx.Request().Input("name"),
	}); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Category created successfully",
		"data":    nil,
	})
}

func (r *CategoryController) Update(ctx http.Context) http.Response {
	var category models.Category
	facades.Orm().Query().Where("id", ctx.Request().Input("id")).FirstOrCreate(&category, models.User{Name: ctx.Request().Input("name")})

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Category updated successfully",
		"data":    nil,
	})
}

func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	var category models.Category
	facades.Orm().Query().Find(&category, ctx.Request().Input("id"))
	res, err := facades.Orm().Query().Delete(&category)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Category deleted successfully",
		"data":    res,
	})
}
