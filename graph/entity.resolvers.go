package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/skill_service/graph/generated"
	"dudo/skill_service/graph/model"
	"fmt"
	"math/rand"
)

func (r *entityResolver) FindTechnologyByID(ctx context.Context, id string) (*model.Technology, error) {
	tech := &model.Technology{
		ID: fmt.Sprintf("T%d", rand.Intn(100)),
	}
	return tech, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
