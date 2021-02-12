package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jrmbrgs/goodkarma-api/library/graph/generated"
	"github.com/jrmbrgs/goodkarma-api/library/models"
)

func (r *mutationResolver) CreateNewCharityOrganizationTag(ctx context.Context, input models.NewCharityOrganizationTag) (*models.CharityOrganizationTag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FetchAllCharityOrganizationTags(ctx context.Context) ([]*models.CharityOrganizationTag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
