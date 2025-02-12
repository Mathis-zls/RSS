package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mathis-zls/RSS/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type param struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := param{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	feedfollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		IDFeed:    params.FeedID,
		IDUser:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Create Feed_Follow failed: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCfg *apiConfig) handlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedfollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Get Feed_Follows failed: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollow(feedfollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollowid := chi.URLParam(r, "feedFollowID")
	id, err := uuid.Parse(feedfollowid)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Parse from UUID ID feedfollow failed: %v", err))
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		IDUser: user.ID,
		ID:     id,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Delete Feed_Follows failed: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
