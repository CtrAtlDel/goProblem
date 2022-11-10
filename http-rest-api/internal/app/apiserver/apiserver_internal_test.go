package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_HandleHello(t *testing.T) {
	s := New(NewConfig())                                    // создаем новый конфиг
	rec := httptest.NewRecorder()                            // делаем новую запись
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) // создаем новый http запрос
	s.handleHello().ServeHTTP(rec, req) 
	assert.Equal(t, rec.Body.String(), "hello")
}
