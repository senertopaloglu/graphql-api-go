package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.74

import (
	"context"
	"fmt"

	main "github.com/senertopaloglu/graphql-api-go"
	"github.com/senertopaloglu/graphql-api-go/graph/generated"
)

// WalletBalance is the resolver for the walletBalance field.
func (r *queryResolver) WalletBalance(ctx context.Context, address string) (*main.Balance, error) {
	panic(fmt.Errorf("not implemented: WalletBalance - walletBalance"))
}

// TokenPrice is the resolver for the tokenPrice field.
func (r *queryResolver) TokenPrice(ctx context.Context, id string) (*main.TokenPrice, error) {
	panic(fmt.Errorf("not implemented: TokenPrice - tokenPrice"))
}

// TransactionHistory is the resolver for the transactionHistory field.
func (r *queryResolver) TransactionHistory(ctx context.Context, address string, limit *int) ([]*main.Transaction, error) {
	panic(fmt.Errorf("not implemented: TransactionHistory - transactionHistory"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
