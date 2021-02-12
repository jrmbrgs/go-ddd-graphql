package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/jrmbrgs/goodkarma-api/library/database"
	"github.com/jrmbrgs/goodkarma-api/library/models"
)

func (r *mutationResolver) CreateNewOffenseCategory(ctx context.Context, input models.NewOffenseCategory) (*models.OffenseCategory, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	fmt.Println("input", input.Label, input.Description)
	offenseCategory := models.OffenseCategory{}
	offenseCategory.Label = input.Label
	offenseCategory.Description = input.Description
	offenseCategory.ParentID = input.ParentID
	db.Create(&offenseCategory)
	return &offenseCategory, nil
}

func (r *queryResolver) FetchAllOffenseCategories(ctx context.Context) ([]*models.OffenseCategory, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}

	var offenseCategories []*models.OffenseCategory
	//var s = "nil"
	//db.Where(&models.OffenseCategory{ParentID: &s}).Find(&offenseCategories)
	db.Where("parent_id is null").Find(&offenseCategories)
	for _, offenseCategory := range offenseCategories {
		var subCategories []*models.OffenseCategory
		db.Where(&models.OffenseCategory{ParentID: offenseCategory.ParentID}).Find(&subCategories)
		//offenseCategory.SubCategories = subCategories
	}
	return offenseCategories, nil
}
