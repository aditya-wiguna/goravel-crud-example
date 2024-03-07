package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	categoryController := controllers.NewCategoryController()

	facades.Route().Get("/users/{id}", userController.Show)

	//Resource route
	facades.Route().Resource("/category", categoryController)
}
