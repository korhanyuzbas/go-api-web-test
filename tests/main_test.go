package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIRoot(t *testing.T) {
	createEnvironmentConfig()
	router := SetupTestRouter()

	w := PerformRequest(
		router,
		"GET",
		"/",
		nil,
	)
	assert.Equal(t, w.Code, 200)

}
