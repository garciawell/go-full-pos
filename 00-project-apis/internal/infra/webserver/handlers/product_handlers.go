package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/garciawell/go-full-pos/apis/internal/dto"
	"github.com/garciawell/go-full-pos/apis/internal/entity"
	"github.com/garciawell/go-full-pos/apis/internal/infra/database"
	entityPkg "github.com/garciawell/go-full-pos/apis/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductInput true "Product requested"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

// Update Product godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Param product body dto.CreateProductInput true "Product requested"
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Get Product godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query string false "Product number"
// @Param limit query string false "Product limit"
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	sort := r.URL.Query().Get("sort")
	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
