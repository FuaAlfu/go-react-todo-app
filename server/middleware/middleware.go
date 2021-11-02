package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo.driver/bson"
	"go.mongodb.org/mongo.driver/bson/primitive"
	"go.mongodb.org/mongo.driver/mongo"
	"go.mongodb.org/mongo.driver/mongo/options"

)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading the .env file..")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)

	client, err := mongo.connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb..")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	playload := getAllTask()
	json.NewEncoder(w).Encode(playload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {}

func TaskComplete(w http.ResponseWriter, r *http.Request) {}

func UndoTask(w http.ResponseWriter, r *http.Request) {}

func DeleteTask(w http.ResponseWriter, r *http.Request) {}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {}
