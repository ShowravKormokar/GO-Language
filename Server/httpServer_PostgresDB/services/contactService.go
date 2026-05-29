package services

import (
	"encoding/json"
	"net/http"
	"server_postgres/database"
	"server_postgres/models"
	"server_postgres/types"

	"github.com/gorilla/mux"
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

func GetContact(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	var contact models.Contact
	err := database.DB.QueryRow("SELECT id, name, phone, description, created_at FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Description, &contact.CreatedAt)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(types.ErrorResponse{
			Mssg:   "Contact Not found",
			Status: http.StatusNotFound,
			Error:  err,
		})
		return
	}
	json.NewEncoder(rw).Encode(types.APIResponse{
		Mssg:   "Contact found",
		Status: http.StatusOK,
		Data:   contact,
	})
}
