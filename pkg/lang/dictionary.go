package lang

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// A handler for the mongodb language collection.
type Dictionary struct {
	collection *mongo.Collection
}

// Creates a new dictionary handler.
func NewDictionary(collection *mongo.Collection) Dictionary {
	return Dictionary{collection: collection}
}

// Gets a random word. Only words with a description are returned.
func (d *Dictionary) RandWord(num int) ([]string, error) {
	pipeline := []bson.M{{"$match": bson.M{"description": bson.M{"$ne": ""}}}, {"$sample": bson.M{"size": num}}}
	cursor, err := d.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	var words []string

	for cursor.Next(context.Background()) {
		var word Word
		err := cursor.Decode(&word)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		words = append(words, word.Inner)
	}
	return words, nil
}

// Gets a random word of given type. Only words with a description are returned.
func (d *Dictionary) RandType(num int, wordType WordType) ([]string, error) {
	pipeline := []bson.M{{"$match": bson.M{"type": wordType.String()}}, {"$match": bson.M{"description": bson.M{"$ne": ""}}}, {"$sample": bson.M{"size": num}}}
	cursor, err := d.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	var words []string

	for cursor.Next(context.Background()) {
		var word Word
		err := cursor.Decode(&word)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		words = append(words, word.Inner)
	}
	return words, nil
}

// Finds the given word. Returns nil if not found, and errors if data is missing.
func (d *Dictionary) LookupWord(word string) (*Word, error) {
	var result Word
	searchResult := d.collection.FindOne(context.Background(), bson.M{"word": strings.ToLower(word)})
	fmt.Println(searchResult)
	err := searchResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Inner == "" {
		return nil, fmt.Errorf("Word \"%s\" not found", word)
	}
	return &result, nil
}
