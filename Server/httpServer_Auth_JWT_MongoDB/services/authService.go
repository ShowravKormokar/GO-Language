package services

import (
	"context"
	"encoding/json"
	"fmt"
	"httpServer_JWT_MongoDB/database"
	"httpServer_JWT_MongoDB/dto"
	"httpServer_JWT_MongoDB/models"
	"httpServer_JWT_MongoDB/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterAuthService(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var u dto.RegisterRequest
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
	err = database.UserCollection.FindOne(context.Background(), bson.M{
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
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Somethings wrong",
		})
		return
	}

	fmt.Println("Hash Pass:", hashedPassword)

	user := models.User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	res, err := database.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Registration failed",
		})
		return
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(rw).Encode(dto.BasicResponse{
		Success: true,
		Message: "Registration successfull",
	})
}

func LoginAuthService(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var u dto.LoginRequest
	err := json.NewDecoder(rq.Body).Decode(&u)
	fmt.Println("Body:", u.Email, u.Password)
	if err != nil || u.Email == "" || u.Password == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Invalid body, must fill email and password",
		})
		return
	}

	// Find user
	var user models.User
	err = database.UserCollection.FindOne(context.Background(), bson.M{"email": u.Email}).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Invalid credentials",
		})
		return
	}

	// Match password
	err = utils.CheckPassword(user.Password, u.Password)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: false,
			Message: "Invalid credentials",
		})
		return
	}

	// Generate token

	// Success Login Response
	json.NewEncoder(rw).Encode(dto.DataResponse{
		BasicResponse: dto.BasicResponse{
			Success: true,
			Message: "Login successfull",
		},
		Data: user,
	})
}
