// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package server

import (
	"context"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/logic"
	"bookstore/rpc/check/internal/svc"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *check.Request) (*check.Response, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
