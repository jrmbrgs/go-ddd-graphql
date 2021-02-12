package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/jrmbrgs/goodkarma-api/library/database"
	"github.com/jrmbrgs/goodkarma-api/library/models"
)

func (r *mutationResolver) CreateNewHuman(ctx context.Context, input models.NewHuman) (*models.Human, error) {
	return r.HumanService.BecomeAGoodKarmaMember(input), nil
}

func (r *queryResolver) FetchOneHuman(ctx context.Context, id string) (*models.Human, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}

	var human models.Human
	//var s = "nil"
	db.Where(&models.Human{ID: id}).Preload("offense").Find(&human)
	return &human, nil
}
