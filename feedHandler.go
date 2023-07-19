package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedConfiguration struct {
	FeedID          int    `json:"feed_id"`
	FeedName        string `json:"feed_name"`
	FeedUUID        string `json:"feed_uuid"`
	FileSourceName  string `json:"file_source_name"`
	FeedIndexName   string `json:"feed_index_name"`
	TargetPartner   string `json:"target_partner"`
	CallMinutes     int    `json:"call_minutes"`
	Tags            string `json:"tags"`
}

func getAllFeedConfigurations(c *gin.Context) {
	feedConfigurations := []FeedConfiguration{}

	rows, err := db.Query("SELECT * FROM feed_configurations")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var fc FeedConfiguration
		err := rows.Scan(&fc.FeedID, &fc.FeedName, &fc.FeedUUID, &fc.FileSourceName, &fc.FeedIndexName, &fc.TargetPartner, &fc.CallMinutes, &fc.Tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		feedConfigurations = append(feedConfigurations, fc)
	}

	c.JSON(http.StatusOK, feedConfigurations)
}

func getFeedConfiguration(c *gin.Context) {
	id := c.Param("id")

	var fc FeedConfiguration

	err := db.QueryRow("SELECT * FROM feed_configurations WHERE id = ?", id).
		Scan(&fc.FeedID, &fc.FeedName, &fc.FeedUUID, &fc.FileSourceName, &fc.FeedIndexName, &fc.TargetPartner, &fc.CallMinutes, &fc.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fc)
}

func createFeedConfiguration(c *gin.Context) {
	var fc FeedConfiguration
	if err := c.ShouldBindJSON(&fc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertQuery := `INSERT INTO feed_configurations (feed_name, feed_uuid, file_source_name, feed_index_name, target_partner, call_minutes, tags)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(insertQuery, fc.FeedName, fc.FeedUUID, fc.FileSourceName, fc.FeedIndexName, fc.TargetPartner, fc.CallMinutes, fc.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fc)
}

func updateFeedConfiguration(c *gin.Context) {
	id := c.Param("id")

	var fc FeedConfiguration
	if err := c.ShouldBindJSON(&fc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateQuery := `UPDATE feed_configurations SET feed_name = ?, feed_uuid = ?, file_source_name = ?, feed_index_name = ?, target_partners = ?, call_minutes = ?, tags = ?
		WHERE id = ?`

	_, err := db.Exec(updateQuery, fc.FeedName, fc.FeedUUID, fc.FileSourceName, fc.FeedIndexName, fc.TargetPartner, fc.CallMinutes, fc.Tags, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fc)
}

func deleteFeedConfiguration(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM feed_configurations WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Feed configuration deleted successfully")
}
