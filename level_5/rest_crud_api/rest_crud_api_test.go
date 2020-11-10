package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// Data
	payload := productInventory{
		Product: product{
			ID:    1,
			Code:  "Code 1",
			Name:  "Name 1",
			Price: 15.42,
		},
		Quantity: 1,
	}

	data, marshalError := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/", createProduct).Methods("POST")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Post(ts.URL+"/inventory/", "application/json", strings.NewReader(string(data)))
	defer res.Body.Close()

	assert.NoError(t, err)
	assert.NoError(t, marshalError)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, len(inventory), 1, "they should be equal")
	assert.Equal(t, inventory[0].Product.ID, 1, "they should be equal")

}

func TestFailCreateProduct(t *testing.T) {
	// Data
	payload := productInventory{
		Quantity: 1,
	}
	data, marshalError := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/", createProduct).Methods("POST")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Post(ts.URL+"/inventory/", "application/json", strings.NewReader(string(data)))
	defer res.Body.Close()

	assert.NoError(t, err)
	assert.NoError(t, marshalError)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	assert.Equal(t, len(inventory), 1, "they should be equal")

}

func TestGetOneProduct(t *testing.T) {
	fillInventory()

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", getOneProduct).Methods("GET")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Get(ts.URL + "/inventory/1")
	defer res.Body.Close()

	var newInventory productInventory
	resBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(resBody, &newInventory)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, newInventory, inventory[0], "they should be equal")
}

func TestFailGetOneProduct(t *testing.T) {
	fillInventory()

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", getOneProduct).Methods("GET")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Get(ts.URL + "/inventory/u")
	defer res.Body.Close()

	var newInventory productInventory
	resBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(resBody, &newInventory)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")

}

func fillInventory() {
	inventory = []productInventory{
		{
			Product: product{
				ID:    1,
				Code:  "Code 1",
				Name:  "Name 1",
				Price: 15.42,
			},
			Quantity: 1,
		},
		{
			Product: product{
				ID:    2,
				Code:  "Code 2",
				Name:  "Name 2",
				Price: 1225.55,
			},
			Quantity: 2,
		},
	}
}

func TestGetAllProducts(t *testing.T) {
	fillInventory()

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", getAllProducts).Methods("GET")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Get(ts.URL + "/inventory")
	defer res.Body.Close()

	var newInventory []productInventory
	resBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(resBody, &newInventory)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, newInventory, inventory, "they should be equal")

}

func TestDeleteProduct(t *testing.T) {
	fillInventory()

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")

	// Create the request
	req := httptest.NewRequest("DELETE", "/inventory/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	resBody, err := ioutil.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, "The Product with ID 1 has been deleted successfully", string(resBody))
	assert.Len(t, inventory, 1)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, len(inventory), 1, "they should be equal")

}

func TestUpdateProduct(t *testing.T) {
	var newInventory productInventory
	fillInventory()

	// Define the updated data to be sent to the REST endpoint
	payload := productInventory{
		Product: product{
			ID:    1,
			Code:  "Code",
			Name:  "Name",
			Price: 14,
		},
		Quantity: 15,
	}
	data, err := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")

	// Create the request
	req := httptest.NewRequest("PUT", "/inventory/1", strings.NewReader(string(data)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	resBody, err := ioutil.ReadAll(res.Body)
	unmarshalError := json.Unmarshal(resBody, &newInventory)

	assert.NoError(t, err)
	assert.NoError(t, unmarshalError)
	assert.Equal(t, res.Code, http.StatusOK, "they should be equal")
	assert.Equal(t, newInventory, inventory[0], "they should be equal")
}

func TestFailNoProductUpdateProduct(t *testing.T) {
	var newInventory productInventory
	fillInventory()

	// Define the updated data to be sent to the REST endpoint
	payload := productInventory{
		Quantity: 15,
	}
	data, err := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")

	// Create the request
	req := httptest.NewRequest("PUT", "/inventory/1", strings.NewReader(string(data)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	resBody, err := ioutil.ReadAll(res.Body)
	unmarshalError := json.Unmarshal(resBody, &newInventory)

	assert.NoError(t, err)
	assert.NoError(t, unmarshalError)
	assert.Equal(t, res.Code, http.StatusBadRequest, "they should be equal")
	//assert.Equal(t, newInventory, inventory[0], "they should be equal")
}

func TestFailDifferentIDProductUpdateProduct(t *testing.T) {
	var newInventory productInventory
	fillInventory()

	// Define the updated data to be sent to the REST endpoint
	payload := productInventory{
		Product: product{
			ID:    2,
			Code:  "Code",
			Name:  "Name",
			Price: 14,
		},
		Quantity: 15,
	}
	data, err := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")

	// Create the request
	req := httptest.NewRequest("PUT", "/inventory/1", strings.NewReader(string(data)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	resBody, err := ioutil.ReadAll(res.Body)
	unmarshalError := json.Unmarshal(resBody, &newInventory)

	assert.NoError(t, err)
	assert.NoError(t, unmarshalError)
	assert.Equal(t, res.Code, http.StatusBadRequest, "they should be equal")
}

func TestPatchUpdateProduct(t *testing.T) {
	var newInventory productInventory
	fillInventory()

	// Define the updated data to be sent to the REST endpoint
	payload := productInventory{
		Product: product{
			ID:   1,
			Code: "Code",
		},
		Quantity: 2,
	}
	data, err := json.Marshal(payload)

	expected := productInventory{
		Product: product{
			ID:    1,
			Code:  "Code",
			Name:  "Name 1",
			Price: 0,
		},
		Quantity: 2,
	}

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/patch_inventory/{id}", updatePatchProduct).Methods("PATCH")

	// Create the request
	req := httptest.NewRequest("PATCH", "/patch_inventory/1", strings.NewReader(string(data)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	resBody, err := ioutil.ReadAll(res.Body)
	unmarshalError := json.Unmarshal(resBody, &newInventory)

	assert.NoError(t, err)
	assert.NoError(t, unmarshalError)
	assert.Equal(t, res.Code, http.StatusOK, "they should be equal")
	assert.Equal(t, expected, inventory[0], "they should be equal")
}
