package logic

import (
	"context"

	"github.com/Uzemiu/moment-rpc/internal/svc"
	"github.com/Uzemiu/moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMomentLogic {
	return &QueryMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMomentLogic) QueryMoment(in *pb.QueryMomentReq) (*pb.QueryMomentResp, error) {
	momentModel := l.svcCtx.MomentModel
	moments, err := momentModel.QueryMoment(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.QueryMomentResp{
		Moments: moments,
	}, nil
}
