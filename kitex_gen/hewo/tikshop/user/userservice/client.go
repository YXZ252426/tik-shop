// Code generated by Kitex v0.11.3. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Auth(ctx context.Context, request *user.AuthRequest, callOptions ...callopt.Option) (r *user.AuthResponse, err error)
	AdminAuth(ctx context.Context, request *user.AuthRequest, callOptions ...callopt.Option) (r *user.AuthResponse, err error)
	GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error)
	UpdateUser(ctx context.Context, request *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error)
	Register(ctx context.Context, request *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest, callOptions ...callopt.Option) (r *user.UpdatePasswordResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Auth(ctx context.Context, request *user.AuthRequest, callOptions ...callopt.Option) (r *user.AuthResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Auth(ctx, request)
}

func (p *kUserServiceClient) AdminAuth(ctx context.Context, request *user.AuthRequest, callOptions ...callopt.Option) (r *user.AuthResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AdminAuth(ctx, request)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, request)
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, request *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, request)
}

func (p *kUserServiceClient) Register(ctx context.Context, request *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, request)
}

func (p *kUserServiceClient) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest, callOptions ...callopt.Option) (r *user.UpdatePasswordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePassword(ctx, request)
}