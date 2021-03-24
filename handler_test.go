package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	web "github.com/yvv4git/go-analyse-sonarqube"
)

func TestHandlers(t *testing.T) {
	type getParams struct {
		x string
		y string
	}

	type resultSum struct {
		Sum int32
	}

	testCases := []struct {
		name        string
		params      getParams
		result      int32
		description string
	}{
		{
			name: "Simple: 1+1",
			params: getParams{
				x: "1",
				y: "1",
			},
			result:      2,
			description: "Check sum",
		},
		{
			name: "More: 5+7",
			params: getParams{
				x: "5",
				y: "7",
			},
			result:      12,
			description: "Check sum",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/sum/%s/%s", tc.params.x, tc.params.y)
			request, err := http.NewRequest("GET", url, nil)
			assert.Nil(t, err)

			webServer := web.SetupWebApp()
			result, err := webServer.Test(request)
			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)
			assert.Nil(t, err)
			//t.Log(string(body))

			var resSum resultSum
			err = json.Unmarshal(body, &resSum)

			assert.Equal(t, tc.result, resSum.Sum, tc.description)
		})
	}
}
