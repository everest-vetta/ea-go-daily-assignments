package songs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeADeleteCall(requestUrl string) *httptest.ResponseRecorder {
	r := gin.New()
	r.DELETE("/song/:id", GetHandler)

	req, _ := http.NewRequest("DELETE", requestUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestShouldReturnSuccessfulResponse(t *testing.T) {
	deleteResponse := makeADeleteCall("/song/1")
	fmt.Println(deleteResponse)
	assert.Equal(t, http.StatusOK, deleteResponse.Code)
}

func TestShouldReturnBadRequestForUnavailableSong(t *testing.T) {
	deleteResponse := makeADeleteCall("/song/0")
	fmt.Println(deleteResponse)
	assert.Equal(t, http.StatusBadRequest, deleteResponse.Code)
}
