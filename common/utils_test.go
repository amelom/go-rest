package common

import (
	"log"
	"net/http"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	_, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("app.req.id").(string); !ok {
			t.Errorf("app.req.id not in request context: got %q", val)
		}
		DisplayAppError(w, err, "error", 500)
	})
	log.Println(testHandler)
}

func TestLoadAppConfig(t *testing.T) {
	urlVal = "./testdata/config.json"
	loadAppConfig()
}

func TestInitConfig(t *testing.T) {
	urlVal = "./testdata/config.json"
	initConfig()
}
