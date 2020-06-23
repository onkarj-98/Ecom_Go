package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID primitive.ObjectID `json:"id" bson:"_id", omitempty` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	product_id int `json:"id" bson:"product_id"`

	Name string `json:"name" bson:"product_name"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description" bson:"product_desc"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price int `json:"price" bson:"unit_price"`
}

// Products defines a slice of Product
type Products []*Product

func ConnectDB() *mongo.Collection {

	// set client optons

	clientOptions := options.Client().ApplyURI("mongodb://demo_user:12345@cluster0-shard-00-00-u0p80.mongodb.net:27017,cluster0-shard-00-01-u0p80.mongodb.net:27017,cluster0-shard-00-02-u0p80.mongodb.net:27017/inventory_service?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority")

	// connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println("Connected to MongoDB!")
	// provide access to perticular document ps: this is temp
	collection := client.Database("inventory_service").Collection("products")

	return collection
}

// GetProducts returns all products from the database
func GetProducts() Products {

	//findOptions := options.Find()
	var productList []*Product

	//passing bson.D{{}} as the filter matches all documents in collection

	cur, err := ConnectDB().Find(context.TODO(), bson.D{{}})

	//	collection1 = ConnectDB()
	//	cur, err := collection1.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	// finding multiple documets returns a cursor
	//Iterating through the cursor allows us to decode documents one at  a time

	for cur.Next(context.TODO()) {
		// create a value which single document decoded
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		productList = append(productList, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// close the cursor once finished

	cur.Close(context.TODO())
	//fmt.Printf("Found", productList)

	return productList

}

// GetProductByID returns a single product which matches the id from the
// database.
//
//If a product is not found this function returns a ProductNotFound error
// func GetProductByID(id int) (*Product, error) {
// 	i := findIndexByProductID(id)
// 	if id == -1 {
// 		return nil, ErrProductNotFound
// 	}

// 	return productList[i], nil
// }

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
// func UpdateProduct(p Product) error {
// 	i := findIndexByProductID(p.ID)
// 	if i == -1 {
// 		return ErrProductNotFound
// 	}

// 	// update the product in the DB
// 	productList[i] = &p

// 	return nil
// }

// AddProduct adds a new product to the database
// func AddProduct(p Product) {
// 	// get the next id in sequence
// 	maxID := productList[len(productList)-1].ID
// 	p.ID = maxID + 1
// 	productList = append(productList, &p)
// }

// DeleteProduct deletes a product from the database
// func DeleteProduct(id int) error {
// 	i := findIndexByProductID(id)
// 	if i == -1 {
// 		return ErrProductNotFound
// 	}

// 	productList = append(productList[:i], productList[i+1])

// 	return nil
// }

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
// func findIndexByProductID(id int) int {
// 	for i, p := range productList {
// 		if p.ID == id {
// 			return i
// 		}
// 	}

// 	return -1
// }

// var productList = []*Product{
// 	&Product{
// 		ID:          1,
// 		Name:        "Latte",
// 		Description: "Frothy milky coffee",
// 		Price:       2.45,
// 		SKU:         "abc323",
// 	},
// 	&Product{
// 		ID:          2,
// 		Name:        "Esspresso",
// 		Description: "Short and strong coffee without milk",
// 		Price:       1.99,
// 		SKU:         "fjd34",
// 	},
// }
