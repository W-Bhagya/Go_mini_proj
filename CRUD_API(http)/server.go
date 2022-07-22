package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Player struct {
	Jersey string  `json : "jersey"`
	Name string `json : "name"`
	Information *Information `json : "information"`
}

// struct Information associated with struct Player
type Information struct {
	Runs string `json : "runs"`
	Wickets string `json : "wickets"`
}


var players []Player

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func deletPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range players {

		if item.Jersey == params["jersey"] {
			players = append(players[:index],players[index + 1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	// loop through the players
	for _, item := range players {
		if item.Jersey == params["jersey"] {
			json.NewEncoder(w).Encode(item)
		    return
		}
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var player Player

	_ = json.NewDecoder(r.Body).Decode(&player)
	player.Jersey = strconv.Itoa(rand.Intn(1000))
	players = append(players, player)
	json.NewEncoder(w).Encode(player)
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range players {
		if item.Jersey == params["jersey"] {
			players = append(players[:index], players[index + 1:]...)
			var player Player

			_ = json.NewDecoder(r.Body).Decode(&player)
			player.Jersey = params["jersey"]
			players = append(players, player)
			json.NewEncoder(w).Encode(player)
			return
		}
	}
}


func main() {

	// function inside the mux library
	r := mux.NewRouter()

	
	players = append(players, Player{Jersey: "18", Name: "Virat Kohli", Information : &Information{Runs: "5678", Wickets: "45"}})
	players = append(players, Player{Jersey: "07", Name: "MS Dhoni", Information : &Information{Runs: "9478", Wickets: "53"}})
	
	r.HandleFunc("/players",getPlayers).Methods("GET")
	r.HandleFunc("/players/{jersey}",getPlayer).Methods("GET")
	r.HandleFunc("/players",createPlayer).Methods("POST")
	r.HandleFunc("/players/{jersey}",updatePlayer).Methods("PUT")
	r.HandleFunc("/players/{jersey}",deletPlayer).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	// if server doesn't start
	log.Fatal(http.ListenAndServe(":8000",r))
}