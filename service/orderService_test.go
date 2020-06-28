//
// Copyright (C) 2020 - -
//

package service

import (
	"testing"
)

func TestOrderMap(t *testing.T) {
	items := map[string]float64{
		"MLA0": 25,
		"MLA1": 10,
		"MLA2": 15,
	}
	orderMap(&items)
}
