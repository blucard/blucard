package backend

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (b *backendImpl) GetRecord(c *gin.Context) {
	record, err := b.recordDB.GetRecord(c.Param("uuid"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid": record.UUID,
		"data": record.Data,
	})
}

type recordInput struct {
	Data string `json:"data"`
}

func (b *backendImpl) SetRecord(c *gin.Context) {
	UUID := c.Param("uuid")

	var record recordInput
	if err := c.ShouldBindJSON(&record); err != nil {
		log.Println(errors.WithStack(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := b.recordDB.SetRecord(UUID, record.Data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid": UUID,
		"data": record.Data,
	})
}
