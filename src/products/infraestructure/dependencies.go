package infraestructure

import (
	"Abarrotes/src/products/application"
	"Abarrotes/src/products/infraestructure/controller"
)

func Init() (*controllers.ShowProductController, *controllers.CreateProductController) {
	ps := NewMySQL()

	showProductUseCase := application.NewShowProduct(ps)
	createUseCase := application.NewCreateProduct(ps)

	showProductController := controllers.NewShowProductController(showProductUseCase)
	createProductController := controllers.NewCreateProductController(createUseCase)

	return showProductController, createProductController
}
