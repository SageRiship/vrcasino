/*************************************************************
 * Copyright 2022 @Entain
 * @author Tirumala Guntakrindapalli
 ************************************************************/

package env

import (
	"fmt"

	"github.com/Tirumala6032/util/logging/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New(LogLevel)
	LeaderboardCollection = Db.Collection("Leaderboard")
}

// Initializing Database and making connection to the Database.
func InitDatabaseConnection() {
	// Getting Connection URL from configurations.
	connectionURI := fmt.Sprintf("%s://%s:%d/", Conf.Database.DBDriver, Conf.Database.DBHost, Conf.Database.DBPort)
	Log.WithFields(logrus.Fields{
		"Method": "InitDatabaseConnection",
	}).Println(connectionURI)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"Method": "InitDatabaseConnection",
		}).Panic(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"Method": "InitDatabaseConnection",
		}).Panic(err)
	}

	Log.WithFields(logrus.Fields{
		"Method": "InitDatabaseConnection",
	}).Println("Database Connected !!.........")

	// Storing Database instance into Db (Global variable)
	Db = client.Database(Conf.Database.DBName)

}
