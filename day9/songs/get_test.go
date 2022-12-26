package songs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeAGetCall(requestUrl string) *httptest.ResponseRecorder {
	r := gin.New()
	r.GET("/song/:id", GetHandler)

	req, _ := http.NewRequest("GET", requestUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestSuccessResponseForAvailableSong(t *testing.T) {
	w := makeAGetCall("/song/1")
	assert.Equal(t, http.StatusOK, w.Code)
	resBody, _ := io.ReadAll(w.Body)

	var response Song
	json.Unmarshal(resBody, &response)
	assert.Equal(t, "Song1", response.Title)
}

func TestShouldReturnInvalidIdType(t *testing.T) {
	w := makeAGetCall("/song/1a")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUnavailableSong(t *testing.T) {
	w := makeAGetCall("/song/0")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
