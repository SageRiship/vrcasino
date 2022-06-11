package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SageRiship/vrcasino/leaderboardservice/model"
	service "github.com/SageRiship/vrcasino/leaderboardservice/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddLeaderboard(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var leaderboard model.Leaderboard
	_ = json.NewDecoder(request.Body).Decode(&leaderboard)

	leaderboard.Id = primitive.NewObjectID()
	result, err := service.AddLeaderboardService(leaderboard)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(result)

}

func GetAllLeaderboard(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var leaderboards []model.Leaderboard
	leaderboards, err := service.GetAllLeaderboardService()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(leaderboards)

}

func GetLeaderboardById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("Object Id :", id)
	leaderboard, err := service.GetLeaderboardByIdService(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(leaderboard)
}

func GetLeaderboardByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	name := params["name"]
	leaderboard, err := service.GetLeaderboardByNameService(name)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(leaderboard)
}

func GetLeaderboardByUserId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println("Params in Get :", params)
	userId, _ := strconv.Atoi(params["userid"])

	leaderboard, err := service.GetLeaderboardByUserIdService(userId)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(leaderboard)
}

func DeleteAllLeaderboard(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := service.DeleteAllLeaderboard()
	json.NewEncoder(response).Encode(count)
}

func DeleteLeaderboardById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	count, _ := service.DeleteLeaderboardByIdService(id)
	json.NewEncoder(response).Encode(count)
}

func DeleteLeaderboardByUserId(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	userId, _ := strconv.Atoi(params["userid"])

	count := service.DeleteLeaderboardByUserIdService(userId)
	json.NewEncoder(response).Encode(count)
}

func UpdateLeaderboard(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	userId, _ := strconv.Atoi(params["id"])
	var leaderboard model.Leaderboard
	_ = json.NewDecoder(request.Body).Decode(&leaderboard)

	result, _ := service.UpdateLeaderboard(userId, leaderboard)

	json.NewEncoder(response).Encode(result)

}
