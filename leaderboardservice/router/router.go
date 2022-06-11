package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	controller "github.com/SageRiship/vrcasino/leaderboardservice/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	//Leaderboard Endpoints
	router.HandleFunc("/api/leaderboard", isAuthorized(controller.AddLeaderboard)).Methods("POST")
	router.HandleFunc("/api/leaderboard", isAuthorized(controller.GetAllLeaderboard)).Methods("GET")
	router.HandleFunc("/api/leaderboard/{id}", isAuthorized(controller.GetLeaderboardById)).Methods("GET")
	router.HandleFunc("/api/leaderboard/name/{name}", isAuthorized(controller.GetLeaderboardByName)).Methods("GET")
	router.HandleFunc("/api/leaderboard/userid/{userid}", isAuthorized(controller.GetLeaderboardByUserId)).Methods("GET")
	router.HandleFunc("/api/leaderboard/{id}", isAuthorized(controller.DeleteLeaderboardById)).Methods("DELETE")
	router.HandleFunc("/api/leaderboard/userid/{userid}", isAuthorized(controller.DeleteLeaderboardByUserId)).Methods("DELETE")
	router.HandleFunc("/api/leaderboard", isAuthorized(controller.DeleteAllLeaderboard)).Methods("DELETE")
	router.HandleFunc("/api/leaderboard/{id}", isAuthorized(controller.UpdateLeaderboard)).Methods("PUT")
	return router
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			type ValidationRequest struct {
				TokenString string `json:"tokenstring,omitempty"`
			}

			var validationRequest ValidationRequest
			validationRequest.TokenString = reqToken

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(validationRequest)

			resp, err := http.Post("http://localhost:8200/v1/validateToken", "application/json;charset=UTF-8", reqBodyBytes) // Call to auth server for token validation.
			if err != nil {
				fmt.Printf("An error occured during REST call: %s", err.Error())
				http.Error(w, "Unable to validate token.", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("\nNon-OK HTTP status:", resp.StatusCode)
				http.Error(w, "Invalid token.", resp.StatusCode)
				return
			}

			type ValidationResponse struct {
				Name string `json:"username,omitempty"`
				Role string `json:"role,omitempty"`
			}
			var validationResponse ValidationResponse
			err = json.NewDecoder(resp.Body).Decode((&validationResponse))
			if err != nil {
				log.Printf("Unable to parse validation response %s \n", err.Error())
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			endpoint(w, r)

		} else {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
		}
	}
}
