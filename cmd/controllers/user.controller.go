package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notes-app-api/cmd/db"
	"notes-app-api/cmd/models"
	"notes-app-api/cmd/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hass, er := utils.EncrypPass(user.Password)
	if er != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}

	user.ID = primitive.NewObjectID()
	user.Password = hass

	_, err := db.UsersCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("user", user.UserName)

	var userDB models.User
	err := db.UsersCollection.FindOne(ctx, bson.M{"userName": user.UserName}).Decode(&userDB)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	match := utils.ComparePass(userDB.Password, user.Password)

	if !match {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDB)
}
