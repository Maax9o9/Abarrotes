package application

import (
	"Abarrotes/src/products/domain"
	"Abarrotes/src/products/domain/entities"
	"fmt"
)

type ModifyProduct struct {
	repository domain.ProductRepository
}

func NewModifyProduct(repo domain.ProductRepository) *ModifyProduct {
	return &ModifyProduct{repository: repo}
}

func (uc *ModifyProduct) Execute(productID int, product entities.Product) error {
	fmt.Println("Modificando producto con ID:", productID)

	existingProduct, err := uc.repository.GetByID(productID)
	if err != nil {
		return fmt.Errorf("error al obtener producto: %v", err)
	}

	if existingProduct.ID == 0 {
		return fmt.Errorf("no se encontró el producto con ID %d", productID)
	}

	err = uc.repository.Update(productID, product)
	if err != nil {
		return fmt.Errorf("error al modificar producto: %v", err)
	}

	fmt.Println("Producto modificado correctamente")
	return nil
}
