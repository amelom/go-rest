package common

import (
	"testing"
)

func TestStartUp(t *testing.T) {
	urlVal = "./testdata/config.json"
	StartUp()
}
