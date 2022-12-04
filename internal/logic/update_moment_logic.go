package logic

import (
	"context"
	"errors"
	"github.com/Uzemiu/moment-rpc/internal/model"
	"github.com/Uzemiu/moment-rpc/internal/svc"
	"github.com/Uzemiu/moment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMomentLogic {
	return &UpdateMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMomentLogic) UpdateMoment(in *pb.UpdateMomentReq) (*pb.UpdateMomentResp, error) {
	momentModel := l.svcCtx.MomentModel

	// 不查询原数据，直接更新的方法
	err := momentModel.UpdateMomentByIdAndUserId(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateMomentResp{}, nil

	// 使用查询原数据的方式，在原数据的基础上进行更新
	cid, err := primitive.ObjectIDFromHex(in.CatId)
	if err != nil {
		return nil, model.ErrInvalidObjectId
	}
	old, err := momentModel.FindOne(l.ctx, in.GetId())
	if err != nil || old == nil {
		return nil, model.ErrNotFound
	}
	if in.UserId != "" && in.UserId != old.UserId.Hex() {
		return nil, errors.New("无权限进行此操作")
	}

	old.Title = in.GetTitle()
	old.Text = in.GetText()
	old.Photos = in.GetImageUrls()
	old.CatId = cid
	err = momentModel.Update(l.ctx, old)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMomentResp{}, nil
}
