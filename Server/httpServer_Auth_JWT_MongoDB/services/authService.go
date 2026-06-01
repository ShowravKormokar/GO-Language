package services

import (
	"context"
	"encoding/json"
	"fmt"
	"httpServer_JWT_MongoDB/database"
	"httpServer_JWT_MongoDB/dto"
	"httpServer_JWT_MongoDB/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterAuthService(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowes",
		})
		return
	}

	var u models.User
	err := json.NewDecoder(rq.Body).Decode(&u)
	fmt.Println("Body:", u.Name, u.Email, u.Password)
	if err != nil || u.Email == "" || u.Password == "" || u.Name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Invalid body, must fill name, email and password",
		})
		return
	}

	// Check email already exist or not
	var existingUser models.User
	err = database.UserCollection.FindOne(context.Background(), map[string]interface{}{
		"email": u.Email,
	}).Decode(&existingUser)
	if err == nil {
		rw.WriteHeader(http.StatusConflict)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Email already exist",
		})
		return
	}

	// Hash the pass

	u.CreatedAt = time.Now()
	u.Role = "user"
	res, err := database.UserCollection.InsertOne(context.Background(), u)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Registration failed",
		})
		return
	}

	u.ID = res.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(rw).Encode(dto.BasicResponse{
		Success: true,
		Message: "Registration successfull",
	})
}
