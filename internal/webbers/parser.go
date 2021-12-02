package webbers

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fishykins/gutenberg/pkg/lang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const lineStart = 481 //28
const lineEnd = 502   //973900

// Handles the entire parsing process
func ParseDictionary(databaseName string) ([]lang.Word, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Open assets
	webbbersDictionaryFile, err := os.Open("./assets/webbersDictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer webbbersDictionaryFile.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(webbbersDictionaryFile)
	partitioner := NewPartitioner()

	// Scan each line into the parser
	for i := 0; i < lineEnd; i++ {
		if scanner.Scan() {
			if i >= lineStart {
				partitioner.ReadLine(i+1, scanner.Text())
			}
		}
	}

	// Tell the partitioner that we have finished
	partitioner.FinishedReading(lineEnd)

	rawWords := make([]lang.Word, 0)

	// If no db name provided, simply return the words
	if databaseName == "" {
		for i, region := range partitioner.Regions {
			word, err := BuildWord(&region, i)
			if err == nil {
				word.Language = "english"
				word.Source = "webbers"
				rawWords = append(rawWords, *word)
				//fmt.Println(word.Header(formatting.FormatType_Plain))
			}
		}
		return rawWords, nil
	}

	// From here on out, we are just doing exporting to database stuff.
	words := bson.A{}

	for i, region := range partitioner.Regions {
		word, err := BuildWord(&region, i)
		if err == nil {
			word.Language = "english"
			word.Source = "webbers"
			rawWords = append(rawWords, *word)
			words = append(words, word.IntoBson())
		}
	}

	// Create db connection and handle
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	db := client.Database(databaseName)
	wordCollection := db.Collection("words")

	// Insert words into database
	res, err := wordCollection.InsertMany(ctx, words)

	if err != nil {
		msg := fmt.Sprintf("Error inserting words into database: %s", err.Error())
		log.Fatal(msg)
	}

	fmt.Println("Added", len(res.InsertedIDs), "words to the db", wordCollection.Name())

	return rawWords, nil
}
