package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "testovoe/docs"
)

func (h *handler) Router(r *gin.Engine) {
	api := r.Group("api/v1")
	{
		subs := api.Group("/subscriptions")
		{
			subs.POST("", h.CreateSubscription)
			subs.GET("/all", h.SubscriptionsList)
			subs.GET("/calculate", h.SubscriptionsSum)
			subs.GET("/:user_id", h.SubscriptionByUserId)
			subs.PUT("/:id", h.UpdateSubscription)
			subs.DELETE("/:id", h.DeleteSubscription)
		}

		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
