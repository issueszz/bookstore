// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package checker

import (
	"context"

	"bookstore/rpc/check/check"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Response = check.Response
	Request  = check.Request

	Checker interface {
		Check(ctx context.Context, in *Request) (*Response, error)
	}

	defaultChecker struct {
		cli zrpc.Client
	}
)

func NewChecker(cli zrpc.Client) Checker {
	return &defaultChecker{
		cli: cli,
	}
}

func (m *defaultChecker) Check(ctx context.Context, in *Request) (*Response, error) {
	client := check.NewCheckerClient(m.cli.Conn())
	return client.Check(ctx, in)
}
