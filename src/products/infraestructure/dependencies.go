package infraestructure

import (
	"Abarrotes/src/products/application"
	"Abarrotes/src/products/infraestructure/controller"
	"Abarrotes/src/products/infraestructure/repository"
)

func InitializeControllers() (controllers.ShowProductController, controllers.CreateProductController) {
	repo := repository.NewInMemoryProductRepository()

	showProductUseCase := application.NewShowProduct(repo)
	createUseCase := application.NewCreateProduct(repo)

	showProduct := controllers.NewShowProductController(showProductUseCase)
	createController := controllers.NewCreateProductController(createUseCase)

	return *showProduct, *createController
}
