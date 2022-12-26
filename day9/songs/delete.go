package songs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id type"})
		return
	}

	for index, song := range songs {
		if song.Id == idToInt {
			if index != len(songs)-1 {
				songs[index] = songs[len(songs)-1]
			}
			songs = songs[:len(songs)-1]
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "song deleted successfully"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "song doesnot exist"})
}
