package controller

import (
	"encoding/json"
	"net/http"
	"restful-api/dto/request"
	"restful-api/dto/response"
	"restful-api/service"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res := response.NewApiResponseBuilder()
	var productCreateRequest request.CreateProductRequest
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			res.SetStatus(500).SetMessage("Internal Server Error").SetErrors(r).SetRequestId(uuid.New().String())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res.Build())
			return
		}
	}()

	err := json.NewDecoder(req.Body).Decode(&productCreateRequest)
	if err != nil {
		res.SetStatus(400).SetMessage("Invalid request body").SetErrors(err.Error()).SetRequestId(uuid.New().String())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.Build())
		return
	}

	productServiceResponse := pc.productService.CreateProduct(req.Context(), productCreateRequest)

	if productServiceResponse.Status != 201 {
		w.WriteHeader(productServiceResponse.Status)
		res.SetStatus(productServiceResponse.Status).SetMessage(productServiceResponse.Message).SetErrors(productServiceResponse.Errors).SetRequestId(uuid.New().String())
		json.NewEncoder(w).Encode(res.Build())
		return
	}

	res.SetStatus(201).SetMessage("Product created successfully").SetData(productServiceResponse.Data).SetRequestId(uuid.New().String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res.Build())
}
