package config

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Client, error){
    // create context 10 second
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    // create clientOptions
    clientOptions := options.Client().ApplyURI(uri)

    //connecting with ctx and clientOptions
   client, err := mongo.Connect(ctx,clientOptions)
   // if err not nill, return nill and err
   if err != nil {
    return nil,err
   }
   
    return client, nil
}
