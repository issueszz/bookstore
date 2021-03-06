package logic

import (
	"bookstore/rpc/add/add"
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Adder.Add(l.ctx, &add.Request{
		Book: req.Book,
		Price: req.Price,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddResp{Ok: resp.Ok}, nil
}
