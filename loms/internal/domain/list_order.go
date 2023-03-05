package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (d *domain) ListOrder(ctx context.Context, orderID int64) (*Order, error) {
	order, err := getOrder(orderID)
	if err != nil {
		return nil, errors.Wrap(err, "get order")
	}
	return order, nil
}