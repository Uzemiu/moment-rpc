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

	var respMoments []*pb.Moment
	for _, moment := range moments {
		respMoments = append(respMoments, &pb.Moment{
			Id:        moment.ID.Hex(),
			Title:     moment.Title,
			CatId:     moment.CatId.Hex(),
			Text:      moment.Text,
			ImageUrls: moment.Photos,
			UserId:    moment.UserId.Hex(),
			Likes:     int32(len(moment.UserLikes)),
			CreateAt:  moment.CreateAt.UnixMilli(),
		})
	}

	return &pb.QueryMomentResp{
		Moments: respMoments,
	}, nil
}
