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
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Method not allowed",
			Status: http.StatusMethodNotAllowed,
			Error:  nil,
		})
		return
	}

	var c models.Contact
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil || c.Name == "" || c.Phone == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Invalid body, must fill name and phone",
			Status: http.StatusMethodNotAllowed,
			Error:  err,
		})
		return
	}

	// Insert into DB
	query := `INSERT INTO contacts(name, phone, description, created_at) VALUES($1,$2,$3,NOW()) RETURNING id, created_at`

	// For debugging
	// fmt.Printf("Inserting contact: %s, %s, %s\n", c.Name, c.Phone, c.Description)

	err = database.DB.QueryRow(query, c.Name, c.Phone, c.Description).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Failed to insert",
			Status: http.StatusInternalServerError,
			Error:  err,
		})
		return
	}

	// Success response
	json.NewEncoder(rw).Encode(types.SuccessResponse{
		Mssg:   "Contact Created Successfully.",
		Status: http.StatusOK,
	})
}

func GetContacts(rw http.ResponseWriter, req *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, phone, description, created_at FROM contacts")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Failed to fetch",
			Status: http.StatusInternalServerError,
			Error:  err,
		})
		return
	}

	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Description, &c.CreatedAt)
		contacts = append(contacts, c)
	}
	json.NewEncoder(rw).Encode(types.APIResponse{
		Mssg:   "All contacts fetch successfully.",
		Status: http.StatusOK,
		Data:   contacts,
	})
}
