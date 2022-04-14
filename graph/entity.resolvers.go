package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/skill_service/graph/generated"
	"dudo/skill_service/graph/model"
	"fmt"
)

func (r *entityResolver) FindTechnologyByID(ctx context.Context, id string) (*model.Technology, error) {
	skill := &model.Skill{
		ID:         fmt.Sprintf("S%s", id),
		Name:       "foo",
		Proficient: true,
	}

	tech := &model.Technology{
		Skills: []*model.Skill{skill},
	}
	return tech, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
