package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/go_hello_world/graph/generated"
	"dudo/go_hello_world/graph/model"
	"fmt"
	"math/rand"
)

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	skill := &model.Skill{
		ID:         fmt.Sprintf("T%d", rand.Intn(100)),
		Name:       "foo",
		Proficient: true,
	}
	return []*model.Skill{skill}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
