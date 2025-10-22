package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testovoe/internal/errs"
	"testovoe/internal/handler/dto"
)

// CreateSubscription
// @Summary Creating new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body dto.CreateSubscription true "Body for creating a new subscription"
// @Success 200 {object} dto.ErrorResult
// @Failure 400 {object} dto.ErrorResult
// @Failure 500 {object} dto.ErrorResult
// @Router /subscriptions/ [post]
func (h *handler) CreateSubscription(c *gin.Context) {
	var body dto.CreateSubscription

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.CreateSubscription(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// SubscriptionsList
// @Summary Get all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} dto.BaseResult
// @Failure 400 {object} dto.BaseResult
// @Failure 500 {object} dto.BaseResult
// @Router /subscriptions/all [get]
func (h *handler) SubscriptionsList(c *gin.Context) {
	result, err := h.service.SubscriptionsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"error":  nil,
	})
}

// SubscriptionsSum
// @Summary Get sum of subscriptions
// @Description 'start_date' and 'end_date' queries is required. 'user_id' and 'service_name' is optional
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param start_date query string true "Start of subscription. format: 'YYYY-MM-DD' (required)"
// @Param end_date query string true "End of subscription. format: 'YYYY-MM-DD' (required)"
// @Param user_id query string false "Filter by user id (optional)"
// @Param service_name query string false "Filter by service name (optional)"
// @Success 200 {object} dto.BaseResult
// @Failure 400 {object} dto.BaseResult
// @Failure 500 {object} dto.BaseResult
// @Router /subscriptions/calculate [get]
func (h *handler) SubscriptionsSum(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if startDate == "" || endDate == "" {
		log.Println("start_date or end_date empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  "the start_date and end_date queries should not be empty",
		})
		return
	}

	userId := c.Query("user_id")
	serviceName := c.Query("service_name")
	var status int

	result, err := h.service.SubscriptionsSum(startDate, endDate, userId, serviceName)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidTimeFormat) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		c.JSON(status, gin.H{
			"result": nil,
			"error":  err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"error":  nil,
	})
}

// SubscriptionByUserId
// @Summary Get all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id path string true "filter by user id"
// @Success 200 {object} dto.BaseResult
// @Failure 400 {object} dto.BaseResult
// @Failure 404 {object} dto.BaseResult
// @Failure 500 {object} dto.BaseResult
// @Router /subscriptions/{user_id} [get]
func (h *handler) SubscriptionByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  "param user_id is missing",
		})
		return
	}

	result, err := h.service.SubscriptionByUserId(userId)
	if err != nil {
		var status int
		if errors.Is(err, errs.ErrRecordsNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}

		c.JSON(status, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"error":  nil,
	})
}

// UpdateSubscription
// @Summary Update subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "subscription id"
// @Param subscription body dto.UpdateSubscription true "Body for updating a new subscription"
// @Success 200 {object} dto.BaseResult
// @Failure 400 {object} dto.BaseResult
// @Failure 500 {object} dto.BaseResult
// @Router /subscriptions/{id} [put]
func (h *handler) UpdateSubscription(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  "param id is missing",
		})
		return
	}

	var body dto.UpdateSubscription
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	result, err := h.service.UpdateSubscription(idStr, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": *result,
		"error":  nil,
	})
}

// DeleteSubscription
// @Summary Delete subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "subscription id"
// @Success 200 {object} dto.BaseResult
// @Failure 400 {object} dto.BaseResult
// @Failure 500 {object} dto.BaseResult
// @Router /subscriptions/{id} [delete]
func (h *handler) DeleteSubscription(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "param id is missing",
		})
		return
	}

	if err := h.service.DeleteSubscription(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
