package services

import (
	"context"
	"encoding/json"
	"httpServer_JWT_MongoDB/database"
	"httpServer_JWT_MongoDB/dto"
	"httpServer_JWT_MongoDB/middleware"
	"httpServer_JWT_MongoDB/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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
