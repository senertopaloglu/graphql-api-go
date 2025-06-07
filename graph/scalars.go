package graph

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTime converts time.Time to a GraphQL-compliant format
func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
}

// UnmarshalTime parses input into time.Time
func UnmarshalTime(v interface{}) (time.Time, error) {
	return graphql.UnmarshalTime(v)
}
