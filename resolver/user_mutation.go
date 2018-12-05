package resolver

import (
	"github.com/reinhartf/AR-Backend/model"
	"github.com/reinhartf/AR-Backend/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateUser(ctx context.Context, args *struct {
	Email    string
	Password string
}) (*userResolver, error) {
	user := &model.User{
		Email:     args.Email,
		Password:  args.Password,
		IPAddress: *ctx.Value("requester_ip").(*string),
	}

	user, err := ctx.Value("userService").(*service.UserService).CreateUser(user)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created user : %v", *user)
	return &userResolver{user}, nil
}
