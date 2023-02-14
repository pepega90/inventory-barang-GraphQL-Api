package graph

import (
	"context"
	"inventory_graphql_api/internal/ports"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	BarangMasukRepo  ports.IBarangMasukRepository
	BarangKeluarRepo ports.IBarangMasukKeluarRepository
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type barangResolver struct {
	*Resolver
}

func (b *Resolver) BarangKeluar() BarangKeluarResolver {
	return &barangResolver{b}
}

func buildError(ctx context.Context, err error, code int) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}
