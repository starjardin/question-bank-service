package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"question-bank-service/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

// Insert questions by id
func (db *DB) InsertQuestionById(question model.NewQuestion) *model.Question {
	questionColl := db.client.Database("graphql-mongodb-api-db").Collection("question")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := questionColl.InsertOne(ctx, bson.D{{Key: "title", Value: question.Title}, {Key: "response", Value: question.Response}, {Key: "createdAt", Value: question.CreatedAt}, {Key: "lastModifiedAt", Value: question.LastModifiedAt}})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnQuestion := model.Question{ID: insertedID, Title: question.Title, Response: question.Response, LastModifiedAt: question.LastModifiedAt, CreatedAt: question.CreatedAt}

	return &returnQuestion
}

// find question by id
func (db *DB) FindQuestionById(id string) *model.Question {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	questionColl := db.client.Database("graphql-mongodb-api-db").Collection("question")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := questionColl.FindOne(ctx, bson.M{"_id": ObjectID})

	question := model.Question{ID: id}

	res.Decode(&question)

	return &question
}

// find all questions
func (db *DB) Questions() []*model.Question {
	questionColl := db.client.Database("graphql-mongodb-api-db").Collection("question")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := questionColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(questionColl, "Question")

	var questions []*model.Question

	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		var quest *model.Question

		if err = cur.Decode(&quest); err != nil {
			log.Fatal(err)
		}

		questions = append(questions, quest)

		fmt.Println("sus::::::", sus)
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}

		// question := model.Question{ID: (sus[0].Value().StringValue()), Title: (sus[1].Value().StringValue()), Response: (sus[2].Value().StringValue())}
		// question := model.Question{ID: (string(sus[0].String())), Title: (string(sus[1].String()))}

		// questions = append(questions, &quest)
	}
	return questions
}
