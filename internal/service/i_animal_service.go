package service

import "context"

type IAnimalService interface {
	Action(ctx context.Context) string
}
