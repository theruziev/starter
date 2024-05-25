package info

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Handler()

	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	expected := map[string]any{
		"version": Version,
		"commit":  Commit,
	}

	var got map[string]any
	err = json.NewDecoder(rr.Body).Decode(&got)
	require.NoError(t, err)
	require.EqualValues(t, expected, got)

}
