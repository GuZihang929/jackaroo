// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package server

import (
	"context"

	"job/rpc/internal/logic"
	"job/rpc/internal/svc"
	"job/rpc/rpc"
)

type SJobServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedSJobServer
}

func NewSJobServer(svcCtx *svc.ServiceContext) *SJobServer {
	return &SJobServer{
		svcCtx: svcCtx,
	}
}

func (s *SJobServer) Login(ctx context.Context, in *rpc.SerchJobRequest) (*rpc.SerchJobResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
