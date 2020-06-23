package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This  error is raised when product cannot be found in database
var ErrProductNotFound = fmt.Errorf("Product not found!")

// Product defines the structure for the product
//swagger: model

type Product struct {
	// the id for the product
	// primitive id
	//
	// required: false
	// min: 1
	ID primitive.ObjectID `json:"id" bson:"_id", omitempty`

	//id of the product
	//
	// required: true
	//max length 10000000
	PRODUCT_ID int `json:"id" bson:"product_id"`
	// the name of the product
	//
	// required : true
	//max length 100000
	NAME string `json:"name" bson:"product_name"`
	// the description of the prodiuct
	//
	//required : true
	DESCRIPTION string `json:"description" bson:"product_desc"`
	// price of the product
	//
	// required : true
	// min 0.01
	PRICE int `json:"price" bson:"unit_price"`
}

// Products defines  a slice of Product
type Products []*Product

// Connection to Database and returning a collection object
func ConnectDB() *mongo.Collection {

	// set client optons
	clientOptions := options.Client().ApplyURI("mongodb://demo_user:12345@cluster0-shard-00-00-u0p80.mongodb.net:27017,cluster0-shard-00-01-u0p80.mongodb.net:27017,cluster0-shard-00-02-u0p80.mongodb.net:27017/inventory_service?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority")

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// return collection from db
	collection := client.Database("inventory_service").Collection("products")

	return collection
}

// GetProducts returns all products from the database
func GetProducts() Products {

	//findOptions := options.Find()
	var productList []*Product

	//passing bson.D{{}} as the filter matches all documents in collection
	cur, err := ConnectDB().Find(context.TODO(), bson.D{{}})

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

	return productList
}

//will implement GetProductByID , UpdateProduct, AddProduct, DeleteProduct etc.
