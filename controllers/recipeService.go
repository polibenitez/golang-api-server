package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recipes-api/db"
	"recipes-api/models"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Dbconnect()

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	color.Green("⛁ Getting all recipes...")

	var recipes []*models.Recipe

	collection := client.Database("recipes").Collection("recipe")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var recipe models.Recipe
		err := cursor.Decode(&recipe)
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, &recipe)
	}

	json.NewEncoder(w).Encode(recipes)
}

func GetRecipeById(w http.ResponseWriter, r *http.Request) {
	color.Green("⛁ Getting recipe...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	recipe := []models.Recipe{}
	collection := client.Database("recipes").Collection("recipe")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&recipe)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(recipe)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	color.Green("⛁ Creating recipe...")
	w.Header().Set("Content-Type", "application/json")
	var recipe models.Recipe

	json.NewDecoder(r.Body).Decode(&recipe)
	collection := client.Database("recipes").Collection("recipe")
	result, err := collection.InsertOne(context.TODO(), recipe)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(result.InsertedID)
	fmt.Println("Inserted ->", res)
	json.NewEncoder(w).Encode(recipe)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	color.Green("⛁ Updating recipe...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	newRecipe := []models.Recipe{}
	json.NewDecoder(r.Body).Decode(&newRecipe)
	collection := client.Database("recipes").Collection("recipe")

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{primitive.E{Key: "$set", Value: newRecipe}}

	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated ->", res)
	json.NewEncoder(w).Encode(newRecipe)
}
