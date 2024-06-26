package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Ricky004/microservice-server/model"
	"github.com/Ricky004/microservice-server/repository/order"
)

type Order struct {
	Repo *order.RedisRepo
}

func (h *Order) Create(w http.ResponseWriter, r *http.Request) {
   var body struct {
	 CustomerID uuid.UUID   `json:"customer_id"`
	 LineItems []model.LineItem  `json:"line_items"`
   }

   if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  return
   }

   now := time.Now().UTC()

   order := model.Order{
	OrderID: rand.Uint64(),
	CustomerID: body.CustomerID,
	LineItem: body.LineItems,
	CreatedAt: &now,
   }
   
   err := h.Repo.Insert(r.Context(), order)
   if err != nil {
	  fmt.Println("failed to insert:", err)
	  w.WriteHeader(http.StatusInternalServerError)
	  return
   }

   res, err := json.Marshal(order)
   if err != nil {
	fmt.Println("failed to marshal:", err)
	w.WriteHeader(http.StatusInternalServerError)
	return
   }

   w.Write(res)
   w.WriteHeader(http.StatusCreated)
}

func (h *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of all orders")
}

func (h *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}

func (h *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

func (h *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}