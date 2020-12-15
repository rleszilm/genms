// Generated by protoc-gen-go-gsnms-dal. *DO NOT EDIT*
package dal_users

import (
	users "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/example/users"
)

import (
	"context"
	"errors"

	"github.com/rleszilm/gen_microservice/service"
)

var (
	// ErrUserCollectionMethodImpl is returned when the called method is not implemented.
	ErrUserCollectionMethodImpl = errors.New("UserCollection method is not implemented")
)

// UserCollection is an autogenerated interface that can be used to interact with a collection
// of User objects.
type UserCollection interface {
	service.Service

	UserCollectionReader
	UserCollectionWriter
}

// UserCollectionWriter is an autogenerated interface that can be used to write to a collection
// of User objects.
type UserCollectionWriter interface {
	Upsert(context.Context, *users.User) (*users.User, error)
}

// UserCollectionReader is an autogenerated interface that can be used to query a collection
// of User objects. The queries and their values are taken from the representative proto
// message.
type UserCollectionReader interface {
	All(context.Context) ([]*users.User, error)
	Filter(context.Context, *UserFields) ([]*users.User, error)
	ById(context.Context, int64) ([]*users.User, error)
	ByNameAndDivision(context.Context, string, string) ([]*users.User, error)
	StubOnly(context.Context) ([]*users.User, error)
}

// UserFields is an autogenerated struct that
// can be used in the generic queries against UserCollection.
type UserFields struct {
	Id               *int64   `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Division         *string  `json:"division,omitempty"`
	LifetimeScore    *float64 `json:"lifetime_score,omitempty"`
	LastScore        *float32 `json:"last_score,omitempty"`
	LifetimeWinnings *int64   `json:"lifetime_winnings,omitempty"`
	LastWinnings     *int32   `json:"last_winnings,omitempty"`
}

// UnimplementedUserCollection is an autogenerated implementation of
// dal_users.UserCollection that returns an error when any
// method is called.
type UnimplementedUserCollection struct {
	service.Deps
}

// Upsert implements dal_users.UserCollection.Upsert
func (x *UnimplementedUserCollection) Upsert(_ context.Context, _ *users.User) (*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// Filter implements dal_users.UserCollection.Filter
func (x *UnimplementedUserCollection) Filter(_ context.Context, _ *UserFields) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ById implements dal_users.UserCollection.ById
func (x *UnimplementedUserCollection) ById(ctx context.Context, _ int64) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// ByNameAndDivision implements dal_users.UserCollection.ByNameAndDivision
func (x *UnimplementedUserCollection) ByNameAndDivision(ctx context.Context, _ string, _ string) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}

// StubOnly implements dal_users.UserCollection.StubOnly
func (x *UnimplementedUserCollection) StubOnly(ctx context.Context) ([]*users.User, error) {
	return nil, ErrUserCollectionMethodImpl
}
