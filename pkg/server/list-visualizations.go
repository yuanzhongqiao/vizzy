package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/rbren/vizzy/pkg/files"
)

func listVisualizations(c *gin.Context) {
	projectID := c.Param("projectID")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	s3 := files.GetFileManager()
	projectKey := "projects/" + projectID + "/visualizations/"
	files, err := s3.ListDirectories(projectKey)
	if err != nil {
		logrus.WithError(err).Warning("Failed to list files, assuming empty")
		files = []string{}
	}

	c.JSON(http.StatusOK, gin.H{"ids": files})
}