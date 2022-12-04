package logic

import (
	"context"
	"github.com/Uzemiu/moment-rpc/internal/model"
	"github.com/Uzemiu/moment-rpc/internal/svc"
	"github.com/Uzemiu/moment-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertMomentLogic {
	return &InsertMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertMomentLogic) InsertMoment(in *pb.InsertMomentReq) (resp *pb.InsertMomentResp, err error) {
	// 类型转换
	catId := primitive.ObjectID{}
	if in.CatId != "" {
		catId, err = primitive.ObjectIDFromHex(in.CatId)
		if err != nil {
			return nil, err
		}
	}
	userId, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, model.ErrInvalidObjectId
	}

	// 插入数据
	var ID = primitive.NewObjectID()
	err = l.svcCtx.MomentModel.Insert(l.ctx, &model.Moment{
		ID:       ID,
		CatId:    catId,
		Photos:   in.ImageUrls,
		Status:   1,
		Title:    in.Title,
		Text:     in.Text,
		UserId:   userId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertMomentResp{
		Id: ID.Hex(),
	}, nil
}
