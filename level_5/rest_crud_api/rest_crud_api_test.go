package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/", createProduct).Methods("POST")

	ts := httptest.NewServer(router)
	defer ts.Close()

	payload := productInventory{
		Product: product{
			ID:    1,
			Code:  "Code 1",
			Name:  "Name 1",
			Price: 15.42,
		},
		Quantity: 1,
	}

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)

	}
	res, err := http.Post(ts.URL+"/inventory/", "application/json", strings.NewReader(string(data)))
	defer res.Body.Close()

	//fmt.Println(inventory, len(inventory), payload, strings.NewReader(payload))
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, len(inventory), 1, "they should be equal")
	assert.Equal(t, inventory[0].Product.ID, 1, "they should be equal")

}

func TestGetOneProduct(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", getOneProduct).Methods("GET")

	ts := httptest.NewServer(router)
	defer ts.Close()

	fillInventory()

	res, err := http.Get(ts.URL + "/inventory/1")
	defer res.Body.Close()

	var newInventory productInventory
	resBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(resBody, &newInventory)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, newInventory, inventory[0], "they should be equal")

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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", getAllProducts).Methods("GET")

	fillInventory()

	ts := httptest.NewServer(router)
	defer ts.Close()

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")

	server := httptest.NewServer(router)
	//server := NewServer( ...) // criar os handlers
	req := httptest.NewRequest("DELETE", server.URL+"/inventory/1", nil)
	w := httptest.NewRecorder()
	//server.Handler.ServeHTTP(w, req)
	handler := http.HandlerFunc(deleteProduct)
	handler.ServeHTTP(w, req)
	resp := w.Result()
	_, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")

	// ts := httptest.NewServer(router)
	// defer ts.Close()

	// req := httptest.NewRequest("DELETE", ts.URL+"/inventory/1", nil)

	// q := req.URL.Query()
	// q.Add("id", "1")
	// req.URL.RawQuery = q.Encode()

	// res := httptest.NewRecorder()
	// handler := http.HandlerFunc(deleteProduct)

	// handler.ServeHTTP(res, req)

	// fillInventory()

	// //res, err := http.Get(ts.URL + "/inventory/1")
	// //defer res.Body.Close()

	// var newInventory productInventory
	// var oldInventory productInventory = inventory[0]
	// resBody, err := ioutil.ReadAll(res.Body)
	// json.Unmarshal(resBody, &newInventory)

	// assert.Nil(t, err)
	// assert.Equal(t, res.Code, http.StatusOK, "they should be equal")
	// assert.Equal(t, newInventory, oldInventory, "they should be equal")
	// assert.Equal(t, len(inventory), 1, "they should be equal")

}

func TestUpdateProduct(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")

	ts := httptest.NewServer(router)
	defer ts.Close()

	payload := productInventory{
		Product: product{
			ID:    1,
			Code:  "Code 1",
			Name:  "Name 1",
			Price: 15.42,
		},
		Quantity: 1,
	}

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)

	}

	req := httptest.NewRequest("PUT", ts.URL+"/inventory/1", strings.NewReader(string(data)))

	//q := req.URL.Query()
	//q.Add("id", "1")
	//req.URL.RawQuery = q.Encode()

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(updateProduct)

	handler.ServeHTTP(res, req)

	fillInventory()

	//res, err := http.Get(ts.URL + "/inventory/1")
	//defer res.Body.Close()

	var newInventory productInventory
	var oldInventory productInventory = inventory[0]
	resBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(resBody, &newInventory)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, http.StatusOK, "they should be equal")
	assert.Equal(t, newInventory, oldInventory, "they should be equal")

}
