package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Comment  string             `bson:"comment,omitempty" json:"comment,omitempty"`
	UserId   primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type Moment struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	CatId     primitive.ObjectID   `bson:"catId,omitempty" json:"catId,omitempty"`
	Comments  []Comment            `bson:"comments,omitempty" json:"comments,omitempty"`
	Photos    []string             `bson:"photos,omitempty" json:"photos,omitempty"`
	Status    int32                `bson:"status,omitempty" json:"status,omitempty"`
	Title     string               `bson:"title,omitempty" json:"title,omitempty"`
	Text      string               `bson:"text,omitempty" json:"text,omitempty"`
	UserId    primitive.ObjectID   `bson:"userId,omitempty" json:"userId,omitempty"`
	UserLikes []primitive.ObjectID `bson:"userLikes,omitempty" json:"userLikes,omitempty"`
	UpdateAt  time.Time            `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time            `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
