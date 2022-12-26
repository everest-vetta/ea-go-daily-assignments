package songs

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type SongRequest struct {
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"releasedate"`
}

type Song struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"releasedate"`
}

var songs = []Song{
	{Id: 1, Title: "Song1", Singer: "Singer1", Writer: "Writer1", ReleasedDate: time.Now()},
	{Id: 2, Title: "Song2", Singer: "Singer2", Writer: "Writer2", ReleasedDate: time.Now()},
}

func CreateHandler(ctx *gin.Context) {
	var request SongRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err})
		return
	}

	Id := rand.Int()
	newSong := Song{
		Id:           Id,
		Title:        request.Title,
		Singer:       request.Singer,
		Writer:       request.Writer,
		ReleasedDate: time.Time(request.ReleasedDate),
	}
	songs = append(songs, newSong)
	ctx.JSON(http.StatusCreated, newSong)
}
