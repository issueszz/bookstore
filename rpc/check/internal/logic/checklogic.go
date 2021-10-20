package logic

import (
	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"
	"bookstore/rpc/model"
	"context"

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
	book := model.BookStore{Book: in.Book}

	res := l.svcCtx.Db.Find(&book)

	if res.Error != nil {
		return nil, res.Error
	}

	return &check.Response{Found: true, Price: book.Price}, nil
}
