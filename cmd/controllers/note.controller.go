package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"notes-app-api/cmd/db"
	"notes-app-api/cmd/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func timeNow() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now()) - 5*60*60*1000
}

func GetNotes(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.NotesCollection.Find(ctx, bson.D{})
	if err != nil {
		http.Error(w, "Error getting notes", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var notes []models.Note
	if err = cursor.All(ctx, &notes); err != nil {
		http.Error(w, "Error decoding notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func existNote(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var note models.Note
	err := db.NotesCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&note)

	if err != nil {
		return false
	}

	return true
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	_, err := db.NotesCollection.InsertOne(ctx, note)
	if err != nil {
		http.Error(w, "Error creating note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, er := strconv.Atoi(idStr)
	if er != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// TODO: Crear en caso de que no se haya hecho antes, posibles problemas de internet
	// exist := existNote(id)
	//
	// if !exist {
	//
	// }

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	note.ID = id

	filter := bson.M{"_id": id}
	update := bson.M{"$set": note}

	_, err := db.NotesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, er := primitive.ObjectIDFromHex(idStr)
	if er != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}

	_, err := db.NotesCollection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Error deleting note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Note deleted")
}
