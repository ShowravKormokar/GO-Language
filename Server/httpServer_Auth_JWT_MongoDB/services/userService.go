package services

import (
	"context"
	"encoding/json"
	"httpServer_JWT_MongoDB/database"
	"httpServer_JWT_MongoDB/dto"
	"httpServer_JWT_MongoDB/middleware"
	"httpServer_JWT_MongoDB/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserProfile(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	claims := rq.Context().Value(middleware.ClaimsKey).(jwt.MapClaims)

	userID := claims["user_id"].(string)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(
			dto.BasicResponse{
				Success: false,
				Message: "Invalid user id",
			},
		)
		return
	}

	var user models.User
	err = database.UserCollection.FindOne(context.Background(), bson.M{
		"_id": objID,
	}).Decode(&user)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Something went wrong",
		})
		return
	}

	json.NewEncoder(rw).Encode(dto.DataResponse{
		BasicResponse: dto.BasicResponse{
			Success: true,
			Message: "User profile fetched successfully",
		},
		Data: user,
	})
}

// Admin only
func GetAllUsers(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var u []models.User
	cursor, err := database.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Faild to fetch users",
		})
		return
	}

	for cursor.Next(context.Background()) {
		var p models.User
		cursor.Decode(&p)
		u = append(u, p)
	}
	json.NewEncoder(rw).Encode(dto.DataResponse{
		BasicResponse: dto.BasicResponse{
			Success: true,
			Message: "Fetch all users",
		},
		Data: u,
	})
}

func GetUserByID(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	id := mux.Vars(rq)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)

	var u models.User
	err := database.UserCollection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&u)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	json.NewEncoder(rw).Encode(dto.DataResponse{
		BasicResponse: dto.BasicResponse{
			Success: true,
			Message: "User found",
		},
		Data: u,
	})
}

func UpdateUser(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPut {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	id := mux.Vars(rq)["id"]
	objId, _ := primitive.ObjectIDFromHex(id)

	var user models.User
	err := json.NewDecoder(rq.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"updated_at": time.Now(),
		},
	}

	_, err = database.UserCollection.UpdateOne(context.Background(), bson.M{"_id": objId}, update)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	json.NewEncoder(rw).Encode(dto.BasicResponse{
		Success: true,
		Message: "User updated successfully",
	})
}

func DeleteUser(rw http.ResponseWriter, rq *http.Request) {

}
