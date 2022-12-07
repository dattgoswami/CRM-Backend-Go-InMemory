package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customers struct {
	ID        uint32 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted"`
}

var customersList = []Customers{
	{
		ID:        1367,
		Name:      "Rick",
		Role:      "Software Engineer",
		Email:     "rick@example.com",
		Phone:     "55198989",
		Contacted: true,
	},
	{
		ID:        1222,
		Name:      "John",
		Role:      "Data Analyst",
		Email:     "john@example.com",
		Phone:     "75198888",
		Contacted: false,
	},
	{
		ID:        3243,
		Name:      "Ron",
		Role:      "Data Scientist",
		Email:     "ron@example.com",
		Phone:     "45201787",
		Contacted: true,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := json.Marshal(customersList)
	if err != nil {
		fmt.Printf("error marshaling JSON in get customers: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customersList)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramId := mux.Vars(r)["id"]
	flag := false
	paramIdUint := convertStringToUint(paramId)
	for _, customer := range customersList {
		if customer.ID == uint32(paramIdUint) {
			flag = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with required id was not found!")
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customers
	flag := false
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	buffer := []byte(body)
	if err := json.Unmarshal(buffer, &customer); err != nil {
		fmt.Printf("error unmarshaling in add customer JSON: %v\n", err)
	}
	for _, currentCustomer := range customersList {
		if currentCustomer.ID == customer.ID {
			flag = true
		}
	}
	if !flag {
		customersList = append(customersList, customer)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("New customer was added!")
	} else {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("New customer can not be added as there is already a customer with same id!")
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramId := mux.Vars(r)["id"]
	flag := false
	paramIdUint := convertStringToUint(paramId)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var newCustomer Customers
	buffer := []byte(body)
	if err := json.Unmarshal(buffer, &newCustomer); err != nil {
		fmt.Printf("error unmarshaling in update customer JSON: %v\n", err)
	}
	newCustomer.ID = uint32(paramIdUint)
	for i, currentCustomer := range customersList {
		if currentCustomer.ID == newCustomer.ID {
			flag = true
			customersList[i] = newCustomer
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Customer information was updated!")
			break
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with required id was not found!")
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramId := mux.Vars(r)["id"]
	flag := false
	paramIdUint := convertStringToUint(paramId)
	for i, currentCustomer := range customersList {
		if currentCustomer.ID == uint32(paramIdUint) {
			flag = true
			customersList[i] = customersList[len(customersList)-1]
			customersList = customersList[:len(customersList)-1]
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Customer was deleted!")
			break
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with required id was not found!")
	}
}

func convertStringToUint(paramId string) uint32 {
	paramIdUint, err := strconv.ParseUint(paramId, 0, 32)
	if err != nil {
		panic(err)
	}
	return uint32(paramIdUint)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
