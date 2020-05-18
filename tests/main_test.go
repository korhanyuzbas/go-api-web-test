package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func UTestAPIRoot(t *testing.T) {
	//createEnvironmentConfig()
	_, router := SetupTestRouter()

	w := PerformRequest(
		router,
		"GET",
		"/",
		nil,
	)
	assert.Equal(t, w.Code, 200)

}
