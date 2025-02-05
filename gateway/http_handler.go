package main

import (
	"errors"
	"log"

	"github.com/SmoothWay/commons"
	pb "github.com/SmoothWay/commons/api"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{
		client: client,
	}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity
	if err := commons.ReadJSON(r, &items); err != nil {
		commons.WriteError(w, http.StatusBadRequest, "failed to read json")
		return
	}

	if err := validateItems(items); err != nil {
		commons.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	order, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		rStatus := status.Convert(err)
		if rStatus.Code() != codes.InvalidArgument {
			commons.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		log.Printf("failed to createOrder: %v", err)
		commons.WriteError(w, http.StatusInternalServerError, "internal error")
		return
	}

	commons.WriteJSON(w, http.StatusOK, order)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return errors.New("items must have at least one item")
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("item ID is required")
		}

		if i.Quantity <= 0 {
			return errors.New("quantity is required")
		}
	}

	return nil
}
