package logic

import (
	"context"
	"github.com/xh-polaris/moment-rpc/internal/svc"
	"github.com/xh-polaris/moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMomentLogic {
	return &DeleteMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMomentLogic) DeleteMoment(in *pb.DeleteMomentReq) (resp *pb.DeleteMomentResp, err error) {
	momentModel := l.svcCtx.MomentModel
	err = momentModel.DeleteMomentByIdAndUserId(l.ctx, in.GetId(), in.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteMomentResp{}, nil
}
