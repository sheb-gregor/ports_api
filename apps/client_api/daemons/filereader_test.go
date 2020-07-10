package daemons

import (
	"strings"
	"testing"

	"ports_api/apps/client_api/types"

	"github.com/stretchr/testify/assert"
)

func Test_readPortsStream(t *testing.T) {
	outBus := make(chan portData)

	go func() {
		err := readPortsStream(strings.NewReader(testData()), outBus)
		assert.NoError(t, err)
	}()

	pl := portList()
	plLen := len(pl)

	var counter int
	for {
		if counter >= plLen {
			return
		}

		td := <-outBus
		expected, ok := pl[td.Unlocode]
		assert.True(t, ok)
		assert.Equal(t, expected, td.Port)
		counter++
	}
}

func portList() map[string]types.Port {
	return map[string]types.Port{
		"AEAJM": {
			Name:        "Ajman",
			City:        "Ajman",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: [2]float32{55.5136433, 25.4052165},
			Province:    "Ajman",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEAJM"},
			Code:        "52000",
		},
		"AEAUH": {
			Name:        "Abu Dhabi",
			City:        "Abu Dhabi",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: [2]float32{54.37, 24.47},
			Province:    "Abu Z¸aby [Abu Dhabi]",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEAUH"},
			Code:        "52001",
		},
		"AEDXB": {
			Name:        "Dubai",
			City:        "Dubai",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: [2]float32{55.27, 25.25},
			Province:    "Dubayy [Dubai]",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEDXB"},
			Code:        "52005",
		},
	}
}

func testData() string {
	return `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  }
}`
}
