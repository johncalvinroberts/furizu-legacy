package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johncalvinroberts/furizu/graph/generated"
	"github.com/johncalvinroberts/furizu/graph/model"
	"github.com/johncalvinroberts/furizu/whoami"
)

func (r *mutationResolver) StartWhoamiChallenge(ctx context.Context, email string) (*model.EmptyResponse, error) {
	err := whoami.Start(email)
	return &model.EmptyResponse{
		Success: err == nil,
	}, err
}

func (r *mutationResolver) RedeemWhoamiChallenge(ctx context.Context, email string, token string) (*model.JwtResponse, error) {
	jwt, err := whoami.Redeem(email, token)
	return &model.JwtResponse{
		Jwt:     jwt,
		Success: err == nil,
	}, err
}

func (r *mutationResolver) RefreshJwt(ctx context.Context) (*model.JwtResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RevokeToken(ctx context.Context) (*model.EmptyResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
