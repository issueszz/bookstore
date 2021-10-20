package logic

import (
	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"
	"bookstore/rpc/model"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.Request) (*add.Response, error) {
	// todo: add your logic here and delete this line
	//_, err := l.svcCtx.Model.Insert(model.Book{
	//	Book: in.Book,
	//	Price: in.Price,
	//})
	res := l.svcCtx.Db.Create(&model.BookStore{
		Book: in.Book,
		Price: in.Price,
	})

	if res.Error != nil {
		return nil, res.Error
	}

	return &add.Response{Ok: true}, nil
}
