package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/jrmbrgs/goodkarma-api/library/database"
	"github.com/jrmbrgs/goodkarma-api/library/graph/generated"
	"github.com/jrmbrgs/goodkarma-api/library/models"
)

func (r *mutationResolver) CreateNewOffense(ctx context.Context, input models.NewOffense) (*models.Offense, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	fmt.Println("input", input.Info)
	offense := models.Offense{}
	//offenseCategory := models.OffenseCategory{}
	offense.HumanID = input.HumanID
	offense.Info = input.Info
	offense.OffenseCategoryID = input.CategoryID
	db.Create(&offense)
	return &offense, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
