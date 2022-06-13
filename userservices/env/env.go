/*************************************************************
 * Copyright 2022 @Entain
 * @author Tirumala Guntakrindapalli
 ************************************************************/

package env

import (
	"context"
	"log"

	"github.com/Tirumala6032/util/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Db                          *mongo.Database
	UserCollection              *mongo.Collection
	WalletCollection            *mongo.Collection
	WalletTransactionCollection *mongo.Collection
	Conf                        *config.Config
	Ctx                         = context.TODO()
	FileType                    = "yaml"
	LogLevel                    string
)

func Init(path string) {
	initConfig(path, FileType)
}

// Initializing Congiguration folder and filename.
func initConfig(path string, fileType string) {
	var err error

	Conf, err = config.InitConfiguration(path, fileType)
	if err != nil {
		log.Fatal(err)
	}

	LogLevel = Conf.LogLevel
}
