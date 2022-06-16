package main

import (
	"context"
	"flag"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDb *mongo.Database
var (
	databaseUri   string
	approvedBots  uint64
	certifiedBots uint64
	deniedBots    uint64
	userId        string
)

func main() {
	flag.StringVar(&databaseUri, "database-uri", "mongodb://127.0.0.1:27017/infinity", "MongoDB URI")
	flag.Uint64Var(&approvedBots, "approved", 0, "Number of approved bots")
	flag.Uint64Var(&certifiedBots, "certified", 0, "Number of certified bots")
	flag.Uint64Var(&deniedBots, "denied", 0, "Number of denied bots")
	flag.StringVar(&userId, "user", "", "User ID")

	flag.Parse()

	if userId == "" {
		panic("User ID is required")
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUri))
	if err != nil {
		panic(err)
	}

	mongoDb = client.Database("infinity")

	// Set staff stats
	mongoDb.Collection("users").UpdateOne(ctx, bson.M{"userID": userId}, bson.M{"$set": bson.M{"staff_stats": bson.M{"approved_bots": approvedBots, "certified_bots": certifiedBots, "denied_bots": deniedBots}}})
	mongoDb.Collection("users").UpdateOne(ctx, bson.M{"userID": userId}, bson.M{"$set": bson.M{"new_staff_stats": bson.M{"approved_bots": approvedBots, "certified_bots": certifiedBots, "denied_bots": deniedBots}}})
}
