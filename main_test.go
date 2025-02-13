package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorretRecuest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, list, 4)
}

func TestMainHandlerInvalidCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moskov", nil) // искаверкал название города в запросе, иммитация опечатки

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}
