package streaming

import (
	"net/http/httptest"
	"testing"

	"github.com/go-ozzo/ozzo-routing"
)

func TestGet(test *testing.T) {
	ozzoContext := routing.NewContext(httptest.NewRecorder(), nil)
	Get(ozzoContext)
}

func TestStreamDatet(test *testing.T) {
	ozzoContext := routing.NewContext(httptest.NewRecorder(), nil)
	StreamDate(ozzoContext)
}