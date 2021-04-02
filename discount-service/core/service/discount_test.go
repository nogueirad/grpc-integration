package service

import (
	"discount-service/core/domain"
	"discount-service/core/repository"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		userRepository    repository.UserRepository
		productRepository repository.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want DiscountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.userRepository, tt.args.productRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_discountService_GetDiscount(t *testing.T) {
	type fields struct {
		userRepository    repository.UserRepository
		productRepository repository.ProductRepository
	}
	type args struct {
		productId string
		userId    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Discount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &discountService{
				userRepository:    tt.fields.userRepository,
				productRepository: tt.fields.productRepository,
			}
			got, err := c.GetDiscount(tt.args.productId, tt.args.userId)
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