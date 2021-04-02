package handler

import (
	"context"
	"discount-service/core/service"
)

type discountGrpcHandler struct {
	service service.DiscountService
}

func NewDiscountGrpcHandler(service service.DiscountService) *discountGrpcHandler {
	return &discountGrpcHandler{service}
}

func (s *discountGrpcHandler) GetDiscount(context.Context, *DiscountRequest) (*DiscountReply, error) {
	response := &DiscountReply{
		Discount: "25",
	}
	return response, nil
}
