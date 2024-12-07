package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/maxwelbm/pod_example/internal/model"
	"github.com/maxwelbm/pod_example/internal/service"
	pkgErr "github.com/maxwelbm/pod_example/pkg/error"
)

type RequestBodyProduct struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data    *Data  `json:"data,omitempty"`
	Error   bool   `json:"error"`
}

type Data struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

type ControllerProduct struct {
	ServiceProducts service.Service
}

func (c *ControllerProduct) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBodyProduct
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	product := model.Product{
		Name:         reqBody.Name,
		Quantity:     reqBody.Quantity,
		Code_value:   reqBody.Code_value,
		Is_published: reqBody.Is_published,
		Expiration:   reqBody.Expiration,
		Price:        reqBody.Price,
	}

	productServ, err := c.ServiceProducts.Create(product)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to create product")
		return
	}

	dt := Data{
		ID:           productServ.ID,
		Name:         productServ.Name,
		Code_value:   productServ.Code_value,
		Is_published: productServ.Is_published,
		Expiration:   productServ.Expiration,
		Quantity:     productServ.Quantity,
		Price:        productServ.Price,
	}

	body := &ResponseBodyProduct{
		Message: "Product created",
		Data:    &dt,
		Error:   false,
	}
	respondJSON(w, http.StatusCreated, body)
}

func (c *ControllerProduct) GetAll(w http.ResponseWriter, r *http.Request) {
	productsServ, err := c.ServiceProducts.GetAll()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	respondJSON(w, http.StatusOK, productsServ)
}

func (c *ControllerProduct) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	product, err := c.ServiceProducts.GetID(id)
	if err != nil {
		if errors.Is(err, pkgErr.ErrorNotFound) {
			handleError(w, http.StatusNotFound, err.Error())
			return
		}
		handleError(w, http.StatusInternalServerError, "Failed to retrieve product")
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func (c *ControllerProduct) Search(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid price format")
		return
	}

	products, err := c.ServiceProducts.GetSearch(price)
	if err != nil {
		if errors.Is(err, pkgErr.ErrorNotFound) {
			handleError(w, http.StatusNotFound, err.Error())
			return
		}
		handleError(w, http.StatusInternalServerError, "Failed to search products")
		return
	}
	respondJSON(w, http.StatusOK, products)
}

func NewControllerProducts(service service.Service) *ControllerProduct {
	return &ControllerProduct{
		ServiceProducts: service,
	}
}
