package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	coll "github.com/SageRiship/vrcasino/leaderboardservice/env"
	"github.com/SageRiship/vrcasino/leaderboardservice/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddLeaderboardService(leaderboard model.Leaderboard) (*model.Leaderboard, error) {
	leaderboard.Id = primitive.NewObjectID()
	//leaderboard.CreatedOn = primitive.Timestamp{T: uint32(time.Now().Unix())}

	leaderboard.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	inserted, err := coll.LeaderboardCollection.InsertOne(context.Background(), leaderboard)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted 1 User in db with id: ", inserted.InsertedID)
	return &leaderboard, nil

}

func GetAllLeaderboardService() ([]model.Leaderboard, error) {

	var leaderboards []model.Leaderboard
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetLeaderboardByIdService(id primitive.ObjectID) (*model.Leaderboard, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var object model.Leaderboard

	if err := coll.LeaderboardCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetLeaderboardByNameService(name string) (*model.Leaderboard, error) {
	filter := bson.D{{Key: "uname", Value: name}}
	var object model.Leaderboard

	if err := coll.LeaderboardCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func GetLeaderboardByUserIdService(id int) (*model.Leaderboard, error) {
	filter := bson.D{{Key: "user_id", Value: id}}
	var object model.Leaderboard

	if err := coll.LeaderboardCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func DeleteAllLeaderboard() int64 {

	deleteResult, err := coll.LeaderboardCollection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Number of User delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func DeleteLeaderboardByIdService(id primitive.ObjectID) (int, error) {
	result, err := coll.LeaderboardCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return int(result.DeletedCount), nil
}

func DeleteLeaderboardByUserIdService(id int) int {
	result, err := coll.LeaderboardCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "user_id", Value: id},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return int(result.DeletedCount)
}

func UpdateLeaderboard(userId int, leaderboards model.Leaderboard) (*model.Leaderboard, error) {
	var leaderboard model.Leaderboard
	//objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }
	res, err := GetLeaderboardByUserIdService(userId)
	if err != nil {
		log.Panic()
	}
	updateLeaderboardData := updateFilter(leaderboards)
	updateLeaderboardData["created_on"] = res.CreatedOn
	//updateLeaderboardData["updated_on"] = primitive.Timestamp{T: uint32(time.Now().Unix())}
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: updateLeaderboardData}}

	if err := coll.LeaderboardCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&leaderboard); err != nil {
		return nil, err
	}
	log.Println(leaderboard)
	return &leaderboard, nil
}

func updateFilter(leaderboard model.Leaderboard) map[string]interface{} {
	var num map[string]interface{} = make(map[string]interface{})
	if leaderboard.Amount > 0 {
		num["amount"] = leaderboard.Amount
	}
	if leaderboard.UserId != 0 {
		num["user_id"] = leaderboard.UserId
	}
	if leaderboard.Uname != "" {
		num["uname"] = leaderboard.Uname
	}
	if leaderboard.GameId != "" {
		num["game_id"] = leaderboard.GameId
	}
	if leaderboard.GameName != "" {
		num["game_name"] = leaderboard.GameName
	}
	if leaderboard.GameType != "" {
		num["game_type"] = leaderboard.GameType
	}
	if leaderboard.RoomId != "" {
		num["room_id"] = leaderboard.RoomId
	}
	if leaderboard.TableId != "" {
		num["table_id"] = leaderboard.TableId
	}
	if leaderboard.BetStatus != "" {
		num["bet_status"] = leaderboard.BetStatus
	}
	if leaderboard.Currency != "" {
		num["currency"] = leaderboard.Currency
	}
	if leaderboard.DeviceType != "" {
		num["device_type"] = leaderboard.DeviceType
	}
	if leaderboard.DeviceId != "" {
		num["device_id"] = leaderboard.DeviceId
	}
	if leaderboard.Region != "" {
		num["region"] = leaderboard.Region
	}
	if leaderboard.Country != "" {
		num["country"] = leaderboard.Country
	}
	if leaderboard.CreatedBy != "" {
		num["created_by"] = leaderboard.CreatedBy
	}
	if leaderboard.CreatedOn != primitive.NewDateTimeFromTime(time.Now()) {
		num["created_on"] = leaderboard.CreatedOn
	}

	return num
}

func GetAllLeaderboardByCountryService(country string) ([]model.Leaderboard, error) {

	var leaderboards []model.Leaderboard
	filter := bson.D{{Key: "country", Value: country}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetAllLeaderboardByBetStatusService(bet_status string) ([]model.Leaderboard, error) {

	var leaderboards []model.Leaderboard
	filter := bson.D{{Key: "bet_status", Value: bet_status}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetAllLeaderboardByRegionService(region string) ([]model.Leaderboard, error) {

	var leaderboards []model.Leaderboard
	filter := bson.D{{Key: "region", Value: region}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetAllLeaderboardByGameNameService(game_name string) ([]model.Leaderboard, error) {

	var leaderboards []model.Leaderboard
	filter := bson.D{{Key: "game_name", Value: game_name}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetAllLeaderboardByDateService(date1 string, date2 string) ([]model.Leaderboard, error) {

	dateF := strings.Split(date1, "-")
	//userId, _ := strconv.Atoi(params[dateF[0]])
	year, _ := strconv.Atoi(dateF[0])
	month, _ := strconv.Atoi(dateF[1])
	day, _ := strconv.Atoi(dateF[2])
	dateS := strings.Split(date2, "-")
	year2, _ := strconv.Atoi(dateS[0])
	month2, _ := strconv.Atoi(dateS[1])
	day2, _ := strconv.Atoi(dateS[2])
	var leaderboards []model.Leaderboard
	fromDate := primitive.NewDateTimeFromTime(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	toDate := primitive.NewDateTimeFromTime(time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC))
	filter := bson.D{{Key: "created_on", Value: bson.M{
		"$gt": fromDate,
		"$lt": toDate,
	}}}
	//filter := bson.D{{Key: "created_on", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}

func GetAllLeaderboardByDateTimeService(date1 time.Time, date2 time.Time) ([]model.Leaderboard, error) {

	// dateF := strings.Split(date1, "-")
	// //userId, _ := strconv.Atoi(params[dateF[0]])
	// year, err := strconv.Atoi(dateF[0])
	// month, _ := strconv.Atoi(dateF[1])
	// day, _ := strconv.Atoi(dateF[2])
	// dateS := strings.Split(date2, "-")
	// year2, _ := strconv.Atoi(dateS[0])
	// month2, _ := strconv.Atoi(dateS[1])
	// day2, _ := strconv.Atoi(dateS[2])
	var leaderboards []model.Leaderboard
	fromDate := primitive.NewDateTimeFromTime(date1)
	toDate := primitive.NewDateTimeFromTime(date2)
	filter := bson.D{{Key: "created_on", Value: bson.M{
		"$gt": fromDate,
		"$lt": toDate,
	}}}
	//filter := bson.D{{Key: "created_on", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))}}
	cursor, err := coll.LeaderboardCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var leaderboard model.Leaderboard
		cursor.Decode(&leaderboard)
		leaderboards = append(leaderboards, leaderboard)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return leaderboards, nil

}
