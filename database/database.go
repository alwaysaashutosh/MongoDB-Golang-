package database

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var dbObject *mongo.Database

type Database struct {
	Driver    string `yaml:"driver"`
	DbName    string `yaml:"dbname"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	IdleConns int    `yaml:"idleConns"`
	OpenConns int    `yaml:"openConns"`
	Schema    string `yaml:"schema"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func (config *Database) DbInstance() *mongo.Database {
	once.Do(func() {
		// // credential := viper.GetString("mongodb.credential")
		// var username, password string
		// // if credential == "local" {
		// // username = viper.GetString("mongoDb.user")
		// // password = viper.GetString("mongoDb.password")
		// username = "john"
		// password = "john"
		// uri := viper.GetString("mongoDb.host")
		// uri := "mongodb+srv://mongoakp.uoj6jcz.mongodb.net/?retryWrites=true&w=majority&appName=mongoAkp"
		// database := "jwtvalidation"
		credentials := options.Credential{
			Username: config.Username,
			Password: config.Password,
		}

		log.Debug().Msgf("Connecting to mongoDB\nURI: %s", config.Host)
		clientOptions := options.Client().ApplyURI(config.Host).SetAuth(credentials) //For authentication add: .SetAuth(credentials)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Panic().Msgf("Error Connecting to mongoDB %s", err)
		}
		log.Debug().Msgf("Ping Database: ")
		if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
			log.Panic().Msgf("Error Connecting to mongoDB %s", err)
		}
		log.Info().Msgf("Connected Successfully")
		dbObject = client.Database(config.DbName)
	})
	return dbObject
}

// func abc(){
// 	dbObject.Collection()
// }
