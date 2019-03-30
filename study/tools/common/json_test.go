package common

import (
	"testing"
	"net/http"
	"io"
	"net/http/httptest"
)

const res string = `
{
	"status": 200,
	"message": "OK",
	"data": {
		"id":1
	}
}
`

type Item struct {
	Id int `json:"id"`
}

func TestToResult(t *testing.T) {
	response := mockResponse()
	result, err := ToResult(response)
	if err != nil {
		t.Error(err)
	}
	if result.Status != 200 {
		t.Fatal(result)
	}
}

func TestToData(t *testing.T) {
	response := mockResponse()
	result, _ := ToResult(response)
	item := Item{}
	ToData(result, &item)
	if item.Id != 1 {
		t.Fatal(item)
	}
}

func mockResponse() *http.Response {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, res)
	}
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	return resp
}
