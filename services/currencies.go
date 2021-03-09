package services

import (
	"context"

	"payment-service/models"
	"payment-service/repositories"
)

func NewCurrenciesService(cr repositories.Currencies) Currencies {
	return &currencyService{currenciesRepository: cr}
}

type currencyService struct {
	currenciesRepository repositories.Currencies
}

func (c *currencyService) AllCurrencies(ctx context.Context) ([]models.Currency, error) {
	return c.currenciesRepository.AllCurrencies(ctx)
}
