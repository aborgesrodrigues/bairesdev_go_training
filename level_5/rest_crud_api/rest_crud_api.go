package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	ID    string  `json:"ID"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type productInventory struct {
	Product  product `json:"Product"`
	Quantity int     `json:"Quantity"`
}

var inventory []productInventory

// func findProductIndex(product product) int {
// 	for i, productInventory := range inventory {
// 		if productInventory.Product.ID == product.ID {
// 			return i
// 		}
// 	}
// 	return -1
// }

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newInventory productInventory

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(err)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newInventory)
	inventory = append(inventory, newInventory)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newInventory)
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]

	for _, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == ProductID {
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
	}
	json.Unmarshal(reqBody, &updatedInventory)

	for i, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == ProductID {

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
			inventory = append(inventory[:i], singleInventoryProduct)
			json.NewEncoder(w).Encode(singleInventoryProduct)
			break
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	ProductID := mux.Vars(r)["id"]

	for i, singleInventoryProduct := range inventory {
		if singleInventoryProduct.Product.ID == ProductID {
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
	router.HandleFunc("/inventory/{id}", updateProduct).Methods("PATCH")
	router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
