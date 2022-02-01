package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johncalvinroberts/furizu/app/graph/generated"
	"github.com/johncalvinroberts/furizu/app/graph/model"
	"github.com/johncalvinroberts/furizu/app/whoami"
)

func (r *mutationResolver) StartWhoamiChallenge(ctx context.Context, email string) (*model.EmptyResponse, error) {
	err := whoami.Start(email)
	return &model.EmptyResponse{
		Success: err == nil,
	}, err
}

func (r *mutationResolver) RedeemWhoamiChallenge(ctx context.Context, email string, token string) (*model.JwtResponse, error) {
	tokenSet, err := whoami.Redeem(email, token)
	return &model.JwtResponse{
		AccessToken:  tokenSet.AccessToken,
		RefreshToken: tokenSet.RefreshToken,
		Success:      err == nil,
	}, err
}

func (r *mutationResolver) RefreshToken(ctx context.Context, prevRefreshToken string) (*model.JwtResponse, error) {
	tokenSet, err := whoami.Refresh(prevRefreshToken)
	return &model.JwtResponse{
		AccessToken:  tokenSet.AccessToken,
		RefreshToken: tokenSet.RefreshToken,
		Success:      err == nil,
	}, err
}

func (r *mutationResolver) RevokeToken(ctx context.Context) (*model.EmptyResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := whoami.Me(ctx)
	// TODO: some kind of serialization layer
	// db/repository types should be different from public-facing "model" types
	// can also cut down on this kind of stupid mapping
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:           user.ID,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt.String(),
		LastUpsertAt: user.LastUpsertAt.String(),
	}, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
