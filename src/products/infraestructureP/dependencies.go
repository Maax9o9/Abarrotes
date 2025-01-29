package infraestructure

import (
	"Abarrotes/src/products/application"
	"Abarrotes/src/products/infraestructureP/controller"
)

func Init() (*controllers.ShowProductController, *controllers.CreateProductController, *controllers.RemoveProductController, *controllers.ModifyProductController) {
	ps := NewMySQL()

	showProductUseCase := application.NewShowProduct(ps)
	createProductUseCase := application.NewCreateProduct(ps)
	removeProductUseCase := application.NewRemoveProduct(ps)
	modifyUseCase := application.NewModifyProduct(ps)

	showProductController := controllers.NewShowProductController(showProductUseCase)
	createProductController := controllers.NewCreateProductController(createProductUseCase)
	removeProductController := controllers.NewRemoveProductController(removeProductUseCase)
	modifyProductController := controllers.NewModifyProductController(modifyUseCase)

	return showProductController, createProductController, removeProductController, modifyProductController
}
