package handler

import (
	"context"
	"discount-service/core/service"
	"reflect"
	"testing"
)

func TestNewDiscountGrpcHandler(t *testing.T) {
	type args struct {
		service service.DiscountService
	}
	tests := []struct {
		name string
		args args
		want *discountGrpcHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiscountGrpcHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiscountGrpcHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_discountGrpcHandler_GetDiscount(t *testing.T) {
	type fields struct {
		service service.DiscountService
	}
	type args struct {
		in0 context.Context
		in1 *DiscountRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DiscountReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &discountGrpcHandler{
				service: tt.fields.service,
			}
			got, err := s.GetDiscount(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiscount() got = %v, want %v", got, tt.want)
			}
		})
	}
}