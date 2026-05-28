package services

import (
	"encoding/json"
	"net/http"
	"server_postgres/database"
	"server_postgres/models"
	"server_postgres/types"
)

func CreateContact(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(types.APIResponse{
			Mssg:   "Method not allowed",
			Status: http.StatusMethodNotAllowed,
		})
		return
	}

	var body types.UserRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.Name == "" || body.Phone == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(types.APIResponse{
			Mssg:   "Invalid body, must fill name and phone",
			Status: http.StatusMethodNotAllowed,
		})
		return
	}

	// Insert into DB
	var c models.Contact
	query := `INSERT INTO contacts(name, phone, description, created_at) VALUES($1,$2,$3,NOW()) RETURNING id, created_at`
	err = database.DB.QueryRow(query, c.Name, c.Phone, c.Description).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(types.APIResponse{
			Mssg:   "Failed to insert",
			Status: http.StatusInternalServerError,
		})
		return
	}

	// Success response
	json.NewEncoder(rw).Encode(types.APIResponse{
		Mssg:   "Contact Created Successfully.",
		Status: http.StatusOK,
	})
}
