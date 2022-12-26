package songs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id type"})
		return
	}

	for _, song := range songs {
		if song.Id == idToInt {
			ctx.IndentedJSON(http.StatusOK, song)
			return
		}
	}

	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"messaage": "song not found"})
}
