package service

import (
	"discount-service/core/domain"
	"discount-service/core/repository"
	"time"
)

const (
	DiscountBase        = 0.0
	DiscountHappyBirth  = 5.0
	DiscountBlackFriday = 10.0
	DiscountMax         = 10.0

	BlackFridayDay   = 25
	BlackFridayMonth = time.November
)

type DiscountService interface {
	GetDiscount(productId, userId string) (domain.Discount, error)
}

type discountService struct {
	userRepository    repository.UserRepository
	productRepository repository.ProductRepository
}

func NewService(userRepository repository.UserRepository, productRepository repository.ProductRepository) DiscountService {
	return &discountService{userRepository, productRepository}
}

func (c *discountService) GetDiscount(productId, userId string) (domain.Discount, error) {
	totalDiscount := DiscountBase
	today := time.Now()

	user, err := c.userRepository.GetUserById(userId)
	if err != nil {
		return domain.Discount{}, err
	}

	product, err := c.productRepository.GetProductById(productId)
	if err != nil {
		return domain.Discount{}, err
	}

	//isHappyBirth
	if today.Day() == user.DateOfBirth.Day() && today.Month() == user.DateOfBirth.Month() {
		totalDiscount = DiscountHappyBirth
	}

	//isBlackFriday
	if today.Day() == BlackFridayDay && today.Month() == BlackFridayMonth {
		totalDiscount += DiscountBlackFriday
	}

	//biggerThanMaxDiscount
	if totalDiscount > DiscountMax {
		totalDiscount = DiscountMax
	}

	valueInCents := totalDiscount * float64(product.PriceInCents) / 100.0

	discount := domain.Discount{
		Percentage:   totalDiscount,
		ValueInCents: int64(valueInCents),
	}

	return discount, nil
}