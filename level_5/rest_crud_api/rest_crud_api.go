package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type product struct {
	ID    int     `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type productInventory struct {
	Product  product `json:"product"`
	Quantity int     `json:"quantity"`
}

type productInventoryPatch struct {
	Product  *product `json:"product,omitempty"`
	Quantity int      `json:"quantity,omitempty"`
}

var inventory []productInventory

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newInventory productInventory

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}

	error := json.Unmarshal(reqBody, &newInventory)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if newInventory.Product.ID == 0 || newInventory.Product.Code == "" || newInventory.Product.Name == "" {
		json.NewEncoder(w).Encode(newInventory)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inventory = append(inventory, newInventory)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newInventory)
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]

	for _, singleInventoryProduct := range inventory {
		intProductID, _ := strconv.Atoi(ProductID)
		if singleInventoryProduct.Product.ID == int(intProductID) {
			json.NewEncoder(w).Encode(singleInventoryProduct)
			break
		}
	}
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(inventory)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	var updatedInventory productInventory

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}
	error := json.Unmarshal(reqBody, &updatedInventory)
	if error != nil {
		fmt.Fprintf(w, error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if updatedInventory.Product.ID == 0 || updatedInventory.Product.Code == "" || updatedInventory.Product.Name == "" {
		json.NewEncoder(w).Encode(updatedInventory)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intProductID, _ := strconv.Atoi(ProductID)

	if updatedInventory.Product.ID != intProductID {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(updatedInventory)
		return
	}

	for i, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == intProductID {
			inventory[i] = updatedInventory
			json.NewEncoder(w).Encode(singleInventoryProduct)
			break
		}
	}
}

func updatePatchProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	intProductID, _ := strconv.Atoi(ProductID)
	var updatedInventory productInventory

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}
	error := json.Unmarshal(reqBody, &updatedInventory)
	if error != nil {
		fmt.Fprintf(w, error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if updatedInventory.Product.ID == 0 || updatedInventory.Product.Code == "" || updatedInventory.Product.Name == "" {
		json.NewEncoder(w).Encode(updatedInventory)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == int(intProductID) {

			//Check if the attributes of the product was changed
			if updatedInventory.Product.Code != "" {
				singleInventoryProduct.Product.Code = updatedInventory.Product.Code
			}
			if updatedInventory.Product.Name != "" {
				singleInventoryProduct.Product.Name = updatedInventory.Product.Name
			}
			if singleInventoryProduct.Product.Price != 0 {
				singleInventoryProduct.Product.Price = updatedInventory.Product.Price
			}
			if updatedInventory.Quantity != 0 {
				singleInventoryProduct.Quantity = updatedInventory.Quantity
			}
			inventory[i] = singleInventoryProduct
			json.NewEncoder(w).Encode(singleInventoryProduct)
			break
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]
	intProductID, _ := strconv.Atoi(ProductID)

	for i, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == int(intProductID) {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Fprintf(w, "The Product with ID %v has been deleted successfully", ProductID)
			break
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", createProduct).Methods("POST")
	router.HandleFunc("/inventory", getAllProducts).Methods("GET")
	router.HandleFunc("/inventory/{id}", getOneProduct).Methods("GET")
	//router.HandleFunc("/inventory/{id}", updatePatchProduct).Methods("PATCH")
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
