package response

type CreateProductResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    ProductData `json:"data"`
	Errors  interface{} `json:"errors"`
}

type ProductData struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductType string `json:"product_type"`
}
