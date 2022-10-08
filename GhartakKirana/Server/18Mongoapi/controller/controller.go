package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adithyahemanth/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://Learn_Go:gotut@cluster0.rbbtt.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbname = "Netflix"

const collist = "watchlist"

var collection *mongo.Collection

func init() {
	clientoptions := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection success")

	collection = client.Database(dbname).Collection(collist)

	fmt.Print("collection instance is ready", collection)

}

func InsertOneMovie(movie model.Netflix) {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie in dbName", insert.InsertedID)
}

func updateOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count :", result.ModifiedCount)

}

func deleteOne(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	delete, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deletecount :", delete.DeletedCount)

}

func deleteAllMovies() int64 {
	deleted, _ := collection.DeleteMany(context.Background(), bson.D{{}})

	fmt.Println(deleted.DeletedCount, "all these data")

	return deleted.DeletedCount
}

func getAllMovies() []primitive.M {
	value, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var Movies []primitive.M

	for value.Next(context.Background()) {
		var movie bson.M
		err := value.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		Movies = append(Movies, movie)

	}
	return Movies
}

// Creating the controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allmovies := getAllMovies()

	json.NewEncoder(w).Encode(allmovies)
	return
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	InsertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

	return
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOne(params["id"])

	json.NewEncoder(w).Encode(params["id"])

	return
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOne(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)

	return

}
