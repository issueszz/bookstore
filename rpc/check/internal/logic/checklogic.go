package logic

import (
	"context"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.Request) (*check.Response, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Model.FindOne(in.Book)

	if err != nil {
		return nil, err
	}

	return &check.Response{Found: true, Price: res.Price}, nil
}
