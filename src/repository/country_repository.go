package repository

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/model"
)

type CountryRepository interface {
	List(ctx context.Context) []*model.Country
	FindByID(ctx context.Context, id string) (*model.Country, error)
	Create(ctx context.Context, country *model.Country) error
	Update(ctx context.Context, country *model.Country) error
}
