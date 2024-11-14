package routes

import (
	"context"
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
	group := r.Group("/api")

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
}
