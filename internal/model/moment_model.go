package model

import (
	"context"
	"github.com/Uzemiu/moment-rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const MomentCollectionName = "moment"

var _ MomentModel = (*customMomentModel)(nil)

type (
	// MomentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMomentModel.
	MomentModel interface {
		momentModel
		QueryMoment(ctx context.Context, query *pb.QueryMomentReq) ([]*Moment, error)
		// UpdateMomentByIdAndUserId 根据id和userId更新动态（id, userId包含在Req内）
		UpdateMomentByIdAndUserId(ctx context.Context, req *pb.UpdateMomentReq) error
		// DeleteMomentByIdAndUserId 根据id和userId删除动态
		DeleteMomentByIdAndUserId(ctx context.Context, id, userId string) error
	}

	customMomentModel struct {
		*defaultMomentModel
	}
)

// NewMomentModel returns a model for the mongo.
func NewMomentModel(url, db, collection string, c cache.CacheConf) MomentModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customMomentModel{
		defaultMomentModel: newDefaultMomentModel(conn),
	}
}

func (m *customMomentModel) QueryMoment(ctx context.Context, query *pb.QueryMomentReq) ([]*Moment, error) {
	var resp []*Moment

	option := options.FindOptions{}
	if query.Size <= 0 {
		query.Size = 10
	}
	option.SetLimit(int64(query.Size))
	if query.Page > 0 {
		option.SetSkip(int64((query.Page - 1) * query.Size))
	}
	filter := bson.M{}

	err := m.conn.Find(ctx, &resp, filter, &option)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateMomentByIdAndUserId 不查询原数据，直接更新的方法
func (m *customMomentModel) UpdateMomentByIdAndUserId(ctx context.Context, data *pb.UpdateMomentReq) error {
	key := prefixMomentCacheKey + data.Id

	oid, err := primitive.ObjectIDFromHex(data.Id)
	if err != nil {
		return ErrInvalidObjectId
	}
	filter := bson.M{
		"_id": oid,
	}
	// 如果userId为空，则直接根据id删除
	// 否则增加userId条件
	if data.UserId != "" {
		uid, err := primitive.ObjectIDFromHex(data.UserId)
		if err != nil {
			return ErrInvalidObjectId
		}
		filter["userId"] = uid
	}
	cid := primitive.ObjectID{}
	set := bson.M{
		"title":    data.Title,
		"text":     data.Text,
		"photos":   data.ImageUrls,
		"catId":    cid,
		"updateAt": time.Now(),
	}
	if data.CatId != "" {
		cid, err = primitive.ObjectIDFromHex(data.CatId)
		if err != nil {
			return ErrInvalidObjectId
		}
		set["catId"] = cid
	}

	// 更新数据
	update := bson.M{
		"$set": set,
	}

	option := options.UpdateOptions{}
	option.SetUpsert(false)

	_, err = m.conn.UpdateOne(ctx, key, filter, update, &option)
	return err
}

func (m *customMomentModel) DeleteMomentByIdAndUserId(ctx context.Context, id, userId string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}
	key := prefixMomentCacheKey + id

	filter := bson.M{
		"_id": oid,
	}
	// 如果userId为空，则直接根据id删除
	// 否则增加userId条件
	if userId != "" {
		uid, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			return ErrInvalidObjectId
		}
		filter["userId"] = uid
	}
	_, err = m.conn.DeleteOne(ctx, key, filter)
	return err
}
