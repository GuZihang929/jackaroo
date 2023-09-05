// Code generated by goctl. DO NOT EDIT.
// Source: company.proto

package companyclient

import (
	"context"

	"xiangxiang/jackaroo/service/company/rpc/types/company"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRequest   = company.CreateRequest
	CreateResponse  = company.CreateResponse
	DetailRequest   = company.DetailRequest
	DetailResponse  = company.DetailResponse
	ListRequest     = company.ListRequest
	ListResponse    = company.ListResponse
	RefreshRequest  = company.RefreshRequest
	RefreshResponse = company.RefreshResponse
	RemoveRequest   = company.RemoveRequest
	RemoveResponse  = company.RemoveResponse
	UpdateRequest   = company.UpdateRequest
	UpdateResponse  = company.UpdateResponse

	Company interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
		Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
		List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	}

	defaultCompany struct {
		cli zrpc.Client
	}
)

func NewCompany(cli zrpc.Client) Company {
	return &defaultCompany{
		cli: cli,
	}
}

func (m *defaultCompany) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultCompany) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultCompany) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultCompany) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultCompany) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.Refresh(ctx, in, opts...)
}

func (m *defaultCompany) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	client := company.NewCompanyClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}
