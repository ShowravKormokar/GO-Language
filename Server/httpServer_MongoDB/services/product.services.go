package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server_MongoDB/database"
	"server_MongoDB/model"
	"server_MongoDB/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Method not allowed",
			Status: http.StatusMethodNotAllowed,
			Error:  nil,
		})
		return
	}

	var p model.Product
	err := json.NewDecoder(rq.Body).Decode(&p)
	fmt.Println(p.Title, p.Price, p.Qunatity, p.Status, p.Description)
	if err != nil || p.Title == "" || p.Price < 0.0 || p.Qunatity < 0 || p.Status != (true || false) {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Invalid body, must fill title, price, quantity>0 and status true/false.",
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	p.Created_At = time.Now()

	res, err := database.ProductCollection.InsertOne(context.Background(), p)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Faild to create product.",
			Status: http.StatusInternalServerError,
			Error:  err,
		})
		return
	}

	p.ID = res.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(rw).Encode(types.SuccessResponse{
		Mssg:   "Product created",
		Status: http.StatusOK,
	})

}

func GetProducts(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Method not allowed",
			Status: http.StatusMethodNotAllowed,
			Error:  nil,
		})
		return
	}

	cursor, err := database.ProductCollection.Find(context.Background(), bson.M{})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Faild fetch products.",
			Status: http.StatusInternalServerError,
			Error:  err,
		})
		return
	}

	var products []model.Product
	for cursor.Next(context.Background()) {
		var p model.Product
		cursor.Decode(&p)
		products = append(products, p)
	}
	json.NewEncoder(rw).Encode(types.DataResponse{
		Mssg:   "All products",
		Status: http.StatusOK,
		Data:   products,
	})

}
