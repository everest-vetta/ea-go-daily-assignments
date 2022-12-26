package songs

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func makeAUpdateCall(requestUrl string, reqBody string) *httptest.ResponseRecorder {
	r := gin.New()
	r.PUT("/song/:id", UpdateHandler)

	req, _ := http.NewRequest("PUT", requestUrl, strings.NewReader(reqBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestSuccessfulUpdationOfSong(t *testing.T) {
	song := `{"title": "TestTitle1", "singer": "TestSinger1", "writer": "TestWriter1", "releaseddate": "2022-01-01"}`

	w := makeAUpdateCall("/song/1", song)

	assert.Equal(t, http.StatusOK, w.Code)
}
