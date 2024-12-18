package main

import (
	"context"
	"log"

	"github.com/vitingr/common"
	pb "github.com/vitingr/common/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (service *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, payload *pb.CreateOrderRequest) error {
	if len(payload.Items) == 0 {
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(payload.Items)

	log.Print(mergedItems)

	return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _,  item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}