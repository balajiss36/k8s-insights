package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/balajiss36/k8s-insights/errors"
	"github.com/balajiss36/k8s-insights/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type handler struct {
	context context.Context
	client  *mongo.Client
}

func NewHandler(ctx context.Context, clientDB *mongo.Client) *handler {
	return &handler{
		context: ctx,
		client:  clientDB,
	}
}

func (h *handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/api/v1")

	group.GET("/test", h.Test)
	group.GET("/pod-insights", h.GetPodInsights)
	group.POST("/pod-insights", h.PostPodInsights)
}

func (h *handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "space persons",
	})
}

func (h *handler) GetPodInsights(c *gin.Context) {
	db := h.client.Database("k8s-insights")
	filter := bson.M{}

	cursor, err := db.Collection("pod-insights").Find(c, filter)
	log.Info().Msg("fetching pod insights")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errors.NewError(c, err, http.StatusNotFound, "no pod insights found")
		}
		errors.NewError(c, err, http.StatusInternalServerError, "unable to fetch pod insights")
		return
	}
	var podInsights models.PodInsightsResponse
	err = cursor.All(c, &podInsights)
	if err != nil {
		errors.NewError(c, err, http.StatusInternalServerError, "unable to decode pod insights")
	}
	defer cursor.Close(h.context)
}

func (h *handler) PostPodInsights(c *gin.Context) {
	var podInsights models.PodInsightsRequest

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errors.NewError(c, err, http.StatusBadRequest, "unable to read pod insights")
	}

	err = json.Unmarshal(jsonData, &podInsights)
	if err != nil {
		errors.NewError(c, err, http.StatusBadRequest, "unable to decode pod insights")
	}
	collection := h.client.Database("k8s-insights").Collection("pod-insights")
	found, err := h.CheckPodData(podInsights, collection)
	if err != nil {
		errors.NewError(c, err, http.StatusInternalServerError, "unable to update existing pod insights data")
		return
	}
	if found {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "existing pod data update successfully",
		})
		return
	}

	_, err = collection.InsertOne(c, podInsights)
	if err != nil {
		errors.NewError(c, err, http.StatusInternalServerError, "unable to insert pod insights")
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "pod data inserted successfully",
	})
}

func (h *handler) CheckPodData(podInsights models.PodInsightsRequest, collection *mongo.Collection) (bool, error) {
	filter := bson.M{"pod": podInsights.PodName, "namespace": podInsights.Namespace}
	count, err := collection.CountDocuments(h.context, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		mongoResult := collection.FindOneAndUpdate(h.context, filter, bson.M{"$set": podInsights})
		if mongoResult.Err() != nil {
			return false, mongoResult.Err()
		}
		return true, nil
	}
	return false, nil
}
