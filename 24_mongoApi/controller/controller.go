package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"git/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://udaysiddu492:V2Tp9ObrfHXtkt3P@cluster0.hohpwil.mongodb.net/netflix"

const dbName = "netflix"
const colName = "watchlist"

// most important
var collection *mongo.Collection

//connect with mongodb

func init()  {
	//client option
	clientOption:= options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client,err:=mongo.Connect(context.TODO(),clientOption)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Mongo connection success")
	collection =client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance or refference is ready")
}

// MONGODB helpers - in separate files


//insert 1 record to mongodb

func insertOneMovie(movie model.Netflix)  {
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in Db with id :",inserted.InsertedID)
}

//update 1 record 
func updateOneMovie(movieId string)  {
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update:= bson.M{"$set":bson.M{"watched": true}}
	result,err :=collection.UpdateOne(context.Background(),filter,update)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("modified count :",result.ModifiedCount)
}
// delete one movie from Db
func deleteOneMovie(movieId string)  {
	id,_:= primitive.ObjectIDFromHex(movieId)
	filter:= bson.M{"_id":id}
	deleteCount,err:=collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got delete with delete count :",deleteCount)
}
//delete all records from mongoDb
func deleteAllMovie() int64  {
	// filter:= bson.D{{}}
	// collection.DeleteMany(context.Background(),filter)

	//or other case
	deleteResults,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Movies deleted:",deleteResults.DeletedCount)
	return deleteResults.DeletedCount
}
// get all movies from database

func getAllMovies()  []primitive.M {
	cursor,err:=collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		if err = cursor.Decode(&movie);err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}



///Actual controllers 🚀📚🛂

func GetMyAllMovies(w http.ResponseWriter,r *http.Request)  {
	
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func  CreateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix

	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])


}
func DeleteAMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:= mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
