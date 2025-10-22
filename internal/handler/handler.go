package handler

import (
	"github.com/gin-gonic/gin"
	"testovoe/internal/service"
)

type Handler interface {
	CreateSubscription(c *gin.Context)
	SubscriptionsList(c *gin.Context)
	SubscriptionsSum(c *gin.Context)
	SubscriptionByUserId(c *gin.Context)
	UpdateSubscription(c *gin.Context)
	DeleteSubscription(c *gin.Context)
	Router(r *gin.Engine)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}
