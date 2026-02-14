package repository

import (
	"database/sql"
	"fmt"
	"GoApi/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}
	var products []model.Product
	var productObj model.Product
	for rows.Next(){
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Product{}, err
		}
		products = append(products, productObj)
	}
	rows.Close()

	return products, nil
}