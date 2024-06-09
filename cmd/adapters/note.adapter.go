package adapters

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

func CreateNote(w http.ResponseWriter, r *http.Request) {

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)
	note.ID = primitive.NewObjectID()
	note.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	note.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.NotesCollection.InsertOne(ctx, note)
	if err != nil {
		http.Error(w, "Error creating note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
