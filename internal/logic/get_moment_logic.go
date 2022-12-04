package logic

import (
	"context"
	"github.com/Uzemiu/moment-rpc/internal/model"

	"github.com/Uzemiu/moment-rpc/internal/svc"
	"github.com/Uzemiu/moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentLogic {
	return &GetMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMomentLogic) GetMoment(in *pb.GetMomentReq) (*pb.GetMomentResp, error) {
	res, err := l.svcCtx.MomentModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, model.ErrNotFound
	}

	return &pb.GetMomentResp{
		Moment: &pb.Moment{
			Id:        res.ID.Hex(),
			Title:     res.Title,
			CatId:     res.CatId.Hex(),
			Text:      res.Text,
			ImageUrls: res.Photos,
			UserId:    res.UserId.Hex(),
			CreateAt:  res.CreateAt.UnixMilli(),
		},
	}, nil
}
