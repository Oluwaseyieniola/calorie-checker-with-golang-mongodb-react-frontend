package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Oluwaseyi/calorie-tracker-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = openCollection(client, "Calories")
var validate = validator.New()

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry models.Entry

	if err := c.Bind(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		msg := fmt.Sprintf("order items was mot created...")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}

func GetEntries(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSONP(http.StatusOK, entries)

}

func EntrybyID(c *gin.Context) {
	EntryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry bson.M
	if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)

}

func EntriesByIngredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{"ingredient": ingredient})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)

	c.JSON(http.StatusOK, entries)
}

func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry
	if err := c.Bind(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})

		fmt.Println(validationErr)
		return
	}
	result, err := entryCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"dish":       entry.Dish,
			"fat":        entry.Fat,
			"ingredient": entry.Ingredient,
			"calories":   entry.Calories,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func UpdateIngredient(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient
	if err := c.Bind(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := entryCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.E{"$set", bson.D{{"ingredient", ingredient.Ingredients}}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.MatchedCount)
}

func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)

}
