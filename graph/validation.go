package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/deathwofl/wine-reviews/pkg/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Validation(ctx context.Context, v validator.Validation) bool {
	isValid, errors := v.Validate()

	// f field
	// e error
	if !isValid {
		for f, e := range errors {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: e,
				Extensions: map[string]interface{}{
					"field": f,
				},
			})
		}
	}
	return isValid
}
