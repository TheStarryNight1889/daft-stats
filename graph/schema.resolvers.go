package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"daft-stats/graph/generated"
	"daft-stats/graph/model"
	"fmt"
)

func (r *queryResolver) Property(ctx context.Context, id string) (*model.Property, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Properties(ctx context.Context) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
