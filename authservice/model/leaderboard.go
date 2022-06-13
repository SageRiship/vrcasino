package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Leaderboard struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	UserId     int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Uname      string             `json:"uname,omitempty" bson:"uname,omitempty"`
	GameId     string             `json:"game_id,omitempty" bson:"game_id,omitempty"`
	GameName   string             `json:"game_name,omitempty" bson:"game_name,omitempty"`
	GameType   string             `json:"game_type,omitempty" bson:"game_type,omitempty"`
	RoomId     string             `json:"room_id,omitempty" bson:"room_id,omitempty"`
	TableId    string             `json:"table_id,omitempty" bson:"table_id,omitempty"`
	TableName  string             `json:"table_name,omitempty" bson:"table_name,omitempty"`
	BetStatus  string             `json:"bet_status,omitempty" bson:"bet_status,omitempty"`
	Amount     int                `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency   string             `json:"currency,omitempty" bson:"currency,omitempty"`
	DeviceType string             `json:"device_type,omitempty" bson:"device_type,omitempty"`
	DeviceId   string             `json:"device_id,omitempty" bson:"device_id,omitempty"`
	Region     string             `json:"region,omitempty" bson:"region,omitempty"`
	Country    string             `json:"country,omitempty" bson:"country,omitempty"`
	CreatedBy  string             `json:"created_by,omitempty" bson:"created_by,omitempty"`
	//CreatedOn  primitive.Timestamp `json:"created_on,omitempty" bson:"created_on,omitempty"`
	CreatedOn primitive.DateTime `json:"created_on,omitempty" bson:"created_on,omitempty"`
}
