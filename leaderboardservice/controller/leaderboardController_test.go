package controller

import (
	"log"
	"testing"
	"time"

	collection "github.com/SageRiship/vrcasino/leaderboardservice/env"
	"github.com/SageRiship/vrcasino/leaderboardservice/model"
	service "github.com/SageRiship/vrcasino/leaderboardservice/service"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var expectedLeaderboard = model.Leaderboard{
	Id:         primitive.NewObjectID(),
	UserId:     1,
	Uname:      "rushi",
	GameId:     "1",
	GameName:   "europeanrouletepro",
	GameType:   "multiplayer",
	RoomId:     "101",
	TableId:    "101",
	TableName:  "roulet",
	BetStatus:  "won",
	Amount:     101,
	Currency:   "USD",
	DeviceType: "multiplayer",
	DeviceId:   "europeanrouletepro",
	Region:     "HYD",
	Country:    "india",
	CreatedBy:  "Rushikesh",
	CreatedOn:  primitive.NewDateTimeFromTime(time.Now()),
}

func TestAddLeaderboard(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		insertedLeaderboard, err := service.AddLeaderboardService(model.Leaderboard{
			Id:         id,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		})
		assert.Nil(t, err)
		assert.Equal(t, &model.Leaderboard{
			Id:         insertedLeaderboard.Id,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender"}, insertedLeaderboard)
	})

	mt.Run("custom error duplicate", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		insertedLeaderboard, err := service.AddLeaderboardService(model.Leaderboard{})
		log.Println("Inserted Leaderboard details :", insertedLeaderboard)
		assert.Nil(t, insertedLeaderboard)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})

	mt.Run("simple error", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		insertedLeaderboard, err := service.AddLeaderboardService(model.Leaderboard{})

		assert.Nil(t, insertedLeaderboard)
		assert.NotNil(t, err)
	})
}

func TestGetAllLeaderboard(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: id1},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		})

		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: id2},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		first_data := expectedLeaderboard
		first_data.Id = id1

		second_data := first_data
		second_data.Id = id2

		leaderboards, err := service.GetAllLeaderboardService()

		log.Println(leaderboards)
		assert.Nil(t, err)
		assert.Equal(t, []model.Leaderboard{
			first_data,
			second_data,
		}, leaderboards)
	})
}

func TestGetLeaderboardById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")
		expectedLeaderboard := model.Leaderboard{
			Id:         r,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		}))
		leaderboardResponse, err := service.GetLeaderboardByIdService(r)
		log.Println(leaderboardResponse)
		assert.Nil(t, err)
		assert.Equal(t, &expectedLeaderboard, leaderboardResponse)
	})
}

func TestGetLeaderboardByUserId(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")

		expectedLeaderboard := model.Leaderboard{
			Id:         r,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		}))
		userid := expectedLeaderboard.UserId
		leaderboardResponse, err := service.GetLeaderboardByUserIdService(userid)
		log.Println(leaderboardResponse)
		assert.Nil(t, err)
		assert.Equal(t, &expectedLeaderboard, leaderboardResponse)
	})
}

func TestGetLeaderboardByName(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")

		expectedLeaderboard := model.Leaderboard{
			Id:         r,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		}))
		userName := expectedLeaderboard.Uname
		leaderboardResponse, err := service.GetLeaderboardByNameService(userName)
		log.Println(leaderboardResponse)
		assert.Nil(t, err)
		assert.Equal(t, &expectedLeaderboard, leaderboardResponse)
	})
}

func TestDeleteLeaderboardById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6271f19f53d1db5e111512ac")
		expectedLeaderboard := model.Leaderboard{
			Id:         r,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: r},
			{Key: "user_id", Value: expectedLeaderboard.UserId},
			{Key: "uname", Value: expectedLeaderboard.Uname},
			{Key: "game_id", Value: expectedLeaderboard.GameId},
			{Key: "game_name", Value: expectedLeaderboard.GameName},
			{Key: "game_type", Value: expectedLeaderboard.GameType},
			{Key: "room_id", Value: expectedLeaderboard.RoomId},
			{Key: "table_id", Value: expectedLeaderboard.TableId},
			{Key: "table_name", Value: expectedLeaderboard.TableName},
			{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
			{Key: "amount", Value: expectedLeaderboard.Amount},
			{Key: "currency", Value: expectedLeaderboard.Currency},
			{Key: "device_type", Value: expectedLeaderboard.DeviceType},
			{Key: "device_id", Value: expectedLeaderboard.DeviceId},
			{Key: "country", Value: expectedLeaderboard.Country},
			{Key: "region", Value: expectedLeaderboard.Region},
			{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
			{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
		}))
		leaderboardResponse, err := service.DeleteLeaderboardByIdService(r)
		log.Println(leaderboardResponse)
		assert.Nil(t, err)

	})
}

func TestUpdateLeaderboard(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection.LeaderboardCollection = mt.Coll
		r, _ := primitive.ObjectIDFromHex("6268f2a22ed1269b68417d61")
		expectedLeaderboard := model.Leaderboard{
			Id:         r,
			UserId:     1,
			Uname:      "rajender",
			GameId:     "1",
			GameName:   "europeanrouletepro",
			GameType:   "multiplayer",
			RoomId:     "101",
			TableId:    "101",
			TableName:  "roulet",
			BetStatus:  "won",
			Amount:     101,
			Currency:   "USD",
			DeviceType: "multiplayer",
			DeviceId:   "europeanrouletepro",
			Region:     "HYD",
			Country:    "india",
			CreatedBy:  "rajender",
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: r},
				{Key: "user_id", Value: expectedLeaderboard.UserId},
				{Key: "uname", Value: expectedLeaderboard.Uname},
				{Key: "game_id", Value: expectedLeaderboard.GameId},
				{Key: "game_name", Value: expectedLeaderboard.GameName},
				{Key: "game_type", Value: expectedLeaderboard.GameType},
				{Key: "room_id", Value: expectedLeaderboard.RoomId},
				{Key: "table_id", Value: expectedLeaderboard.TableId},
				{Key: "table_name", Value: expectedLeaderboard.TableName},
				{Key: "bet_status", Value: expectedLeaderboard.BetStatus},
				{Key: "amount", Value: expectedLeaderboard.Amount},
				{Key: "currency", Value: expectedLeaderboard.Currency},
				{Key: "device_type", Value: expectedLeaderboard.DeviceType},
				{Key: "device_id", Value: expectedLeaderboard.DeviceId},
				{Key: "country", Value: expectedLeaderboard.Country},
				{Key: "region", Value: expectedLeaderboard.Region},
				{Key: "created_by", Value: expectedLeaderboard.CreatedBy},
				{Key: "created_on", Value: expectedLeaderboard.CreatedOn},
			}},
		})

		updateParams := map[string]interface{}{
			"uname": "charan",
		}
		log.Println(updateParams)

		result, err := service.UpdateLeaderboard(expectedLeaderboard.UserId, expectedLeaderboard)

		log.Println(result)

		assert.Nil(t, err)
		assert.Equal(t, &expectedLeaderboard, result)
	})
}
