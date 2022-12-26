package songs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeACall(reqBody string) *httptest.ResponseRecorder {
	r := gin.New()
	r.POST("/song", CreateHandler)

	//jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/song", strings.NewReader(reqBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestSuccessfulCreationOfSong(t *testing.T) {
	song := `{"title": "TestTitle1", "singer": "TestSinger1", "writer": "TestWriter1", "releaseddate": "2022-01-01"}`
	w := makeACall(song)

	assert.Equal(t, http.StatusCreated, w.Code)
	resBody, _ := io.ReadAll(w.Body)

	var response Song
	json.Unmarshal(resBody, &response)

	assert.NotEmpty(t, response.Id)
	assert.Equal(t, "TestTitle1", response.Title)
}

func TestShouldReturnBadRequest(t *testing.T) {
	song := `{"title": "TestTitle1" "singer": "TestSinger1"  "writer": "TestWriter1"  "releaseddate": "2022-01-01"}`

	w := makeACall(song)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}
