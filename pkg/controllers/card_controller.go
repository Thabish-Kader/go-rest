package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thabish/go-rest/pkg/config"
	"github.com/thabish/go-rest/pkg/models"
	"github.com/thabish/go-rest/pkg/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	cardscollection *mongo.Collection = config.GetCollection("cardscollection")
)

func CreateCard(c *fiber.Ctx) error {
	fmt.Print("Create endpoint hit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	card := models.Card{}

	if err := c.BodyParser(&card); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CardResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
	}

	newCard := models.Card{
		Id: primitive.NewObjectID(),
		Name: card.Name,
		Number: card.Number,
	}

	_,err := cardscollection.InsertOne(ctx, newCard)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CardResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.CardResponse{Status: http.StatusCreated, Message: "success", Data: newCard})
}

func GetCard(c *fiber.Ctx) error {
	fmt.Print("Get endpoint hit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	card := models.Card{}
	defer cancel()

	cardId := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(cardId)


	err := cardscollection.FindOne(ctx, bson.M{"id": id}).Decode(&card)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CardResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.CardResponse{Status: http.StatusCreated, Message: "success", Data: card})

}

func GetAllCards(c *fiber.Ctx) error {
	fmt.Println("Get all cards endpoint hit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cards := []models.Card{}
	defer cancel()

	result, err := cardscollection.Find(ctx, bson.M{})
	fmt.Print(result)
	defer result.Close(ctx)

	for result.Next(ctx) {
		card := models.Card{}
		err := result.Decode(&card)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards,card )
	}


	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CardResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.CardResponse{Status: http.StatusCreated, Message: "success", Data: cards})

}


func DeleteCard(c *fiber.Ctx) error {
	fmt.Println("Delete endpoint hit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cardId := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(cardId)

	result, err := cardscollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CardResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.CardResponse{Status: http.StatusNotFound, Message: "error", Data: "Card not found"})
	}

	return c.Status(http.StatusCreated).JSON(responses.CardResponse{Status: http.StatusCreated, Message: "success", Data: "Card deleted successfully"})

}

func UpdateCard(c *fiber.Ctx) error {
	fmt.Print("Update endpoint hit")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	card := models.Card{}
	defer cancel()

	cardId := c.Params("id")
	id, _ := primitive.ObjectIDFromHex(cardId)

	if err := c.BodyParser(&card); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.CardResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
	}

	updatedCard := bson.M{"name": card.Name, "number": card.Number}

	result,err := cardscollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": updatedCard})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CardResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	if result.MatchedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.CardResponse{Status: http.StatusNotFound, Message: "error", Data: "Card not found"})
	}

	card.Id = id

	return c.Status(http.StatusCreated).JSON(responses.CardResponse{Status: http.StatusCreated, Message: "success", Data: card})

}

