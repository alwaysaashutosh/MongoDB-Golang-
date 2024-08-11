package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Database
}

func (m *MongoDB) InsertElement(data *bson.D) {
	result, err := m.moviesCollection().InsertOne(
		context.TODO(), data)

	if err != nil {
		fmt.Printf("error while inserting the document. error: %v", err)
		return
	}
	fmt.Printf("elements inserted successfully. result: %v\n", result.InsertedID)
}

func (m *MongoDB) InsertElements(datasets []interface{}) {
	result, err := m.moviesCollection().InsertMany(context.TODO(), datasets)
	if err != nil {
		fmt.Printf("error while inserting the document. error: %v", err)
		return
	}
	fmt.Printf("elements inserted successfully. result: %v\n", result.InsertedIDs)
}
func (m *MongoDB) ReadElements(condition, projection interface{}) {
	cursor, err := m.moviesCollection().Find(context.TODO(), condition, options.Find().SetProjection(projection))
	if err != nil {
		fmt.Printf("error while reading all elements in the document. error: %v \n", err)
		return
	}
	defer cursor.Close(context.TODO())

	var value []interface{}
	err = cursor.All(context.TODO(), &value)
	if err != nil {
		fmt.Printf("decoding all the document. error: %v\n", err)
		return
	}

	for key, elements := range value {
		fmt.Printf("ID: %d | element: %v \n", key, elements)
	}
}
func (m *MongoDB) ReadWithBsonM(condition *bson.M) {
}

// func ()

func (m *MongoDB) UpdateElement() {

}

func (m *MongoDB) UpdateElements() {
}

func (m *MongoDB) moviesCollection() *mongo.Collection {
	return m.Collection("movies")
}
