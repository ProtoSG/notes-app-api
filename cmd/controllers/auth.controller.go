package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"notes-app-api/cmd/db"
	"notes-app-api/cmd/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = primitive.NewObjectID()

	_, err := db.UsersCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	filter := bson.M{"userName": user.UserName, "password": user.Password}

	err := db.UsersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
