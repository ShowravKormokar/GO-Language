package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfoRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type APIResponse struct {
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// Handler POST /userInfo
func fetchUserInfo(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(rw).Encode(APIResponse{
			Msg:    "Method not allowed!",
			Status: http.StatusMethodNotAllowed,
		})
		return
	}

	var body UserInfoRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.ID == "" || body.Name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(APIResponse{
			Msg:    "Invalid request body, must include name and salary",
			Status: http.StatusBadRequest,
		})
		return
	}

	// Success Response
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(APIResponse{
		Msg:    "User info received successfully",
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"id":    body.ID,
			"name":  body.Name,
			"email": "abc@xyz.com",
		},
	})

	fmt.Printf("SUCCESS: %s %s -> Status %d\n", req.Method, req.URL.Path, http.StatusOK)

}

// Handler: GET /ping (simple health check)
func pingHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(APIResponse{
		Msg:    "Server is alive",
		Status: http.StatusOK,
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/userInfo", fetchUserInfo)
	mux.HandleFunc("/ping", pingHandler)

	fmt.Println("Server running on post 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Couldn't connect to server: ", err.Error())
	}
}

/*

POST http://localhost:8080/userInfo
Request Body (raw:JSON):
{
    "id":"1234",
    "name":"Showrav"
}

Response Body (JSON):
{
    "msg": "User info received successfully",
    "status": 200,
    "data": {
        "email": "abc@xyz.com",
        "id": "1234",
        "name": "Showrav"
    }
}

*/
