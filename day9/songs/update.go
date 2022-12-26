package songs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id type"})
		return
	}

	var request Song

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	request.Id = idToInt
	i := -1
	for index, song := range songs {
		if song.Id == idToInt {
			i = index
			break
		}
	}

	if i >= 0 && i < len(songs) {
		songs[i] = request
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "song updated successfully"})
		return
	}

	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "song doesnot exist"})
}
