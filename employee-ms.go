package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type Employee struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Age         int    `bson:"age" json:"age"`
	Designation string `bson:"designation" json:"designation"`
}

func main() {

	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("local")
	collection := db.Collection("Employee")

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idParam := r.URL.Path[len("/employees/"):]

		var employee Employee
		err = collection.FindOne(context.Background(), bson.M{"id": idParam}).Decode(&employee)
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		} else if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(employee)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
