package database

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var dbObject *mongo.Database

type Database interface {
	InsertElement(data *bson.D)
	InsertElements(datasets []interface{})
	ReadElements(condition, projection interface{})
	// ReadWithCondition()
	// ReadOne()
	UpdateElement()
	UpdateElements()
}
type DatabaseConfig struct {
	Driver    string `yaml:"driver"`
	DbName    string `yaml:"dbname"`
	Host      string `yaml:"host"`
	IdleConns int    `yaml:"idleConns"`
	OpenConns int    `yaml:"openConns"`
	Schema    string `yaml:"schema"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func NewDBClient(config *DatabaseConfig) Database {

	once.Do(func() {
		log.Debug().Msgf("Connecting to mongoDB \n URI: %s", config.Host)
		clientOptions := options.Client().ApplyURI(config.Host)
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
	return &MongoDB{dbObject}
}

/*
func getClientOptionsWithLogging(uri string) *options.ClientOptions {
	// Create a command monitor for logging
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			cmd, _ := json.Marshal(evt.Command)
			log.Debug().Msgf("MongoDB Query Started: %s, Command: %s, Details: %s", evt.CommandName, evt.Command, string(cmd))
		},
		Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
			log.Debug().Msgf("MongoDB Query Succeeded: %s, Duration: %dms", evt.CommandName, evt.DurationNanos/1e6)
		},
		Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
			log.Error().Msgf("MongoDB Query Failed: %s, Error: %s", evt.CommandName, evt.Failure)
		},
	}

	return options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)
}
*/
