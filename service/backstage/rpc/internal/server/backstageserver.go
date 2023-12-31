// Code generated by goctl. DO NOT EDIT.
// Source: backstage.proto

package server

import (
	"context"

	"backstage/rpc/internal/logic"
	"backstage/rpc/internal/svc"
	"backstage/rpc/rpc"
)

type BackstageServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedBackstageServer
}

func NewBackstageServer(svcCtx *svc.ServiceContext) *BackstageServer {
	return &BackstageServer{
		svcCtx: svcCtx,
	}
}

func (s *BackstageServer) Get24P(ctx context.Context, in *rpc.PeoRequest) (*rpc.PeoResponse, error) {
	l := logic.NewGet24pLogic(ctx, s.svcCtx)
	return l.Get24P(in)
}

func (s *BackstageServer) GetSeaJob(ctx context.Context, in *rpc.JobRequest) (*rpc.JobResponse, error) {
	l := logic.NewGetSeaJobLogic(ctx, s.svcCtx)
	return l.GetSeaJob(in)
}
