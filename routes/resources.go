package routes

import (
	"fmt"
	"go-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() {
	fmt.Println("Reoutes....")
	r := mux.NewRouter()

	r.HandleFunc("/api/recipes", controllers.GetRecipes).Methods("GET")
	r.HandleFunc("/api/recipes/{id}", controllers.GetRecipeById).Methods("GET")
	r.HandleFunc("/api/recipes", controllers.CreateRecipe).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", controllers.UpdateRecipe).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9999",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
