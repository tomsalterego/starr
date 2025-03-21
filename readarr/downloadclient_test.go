package readarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

const downloadClientResponseBody = `{
    "enable": true,
    "protocol": "torrent",
    "priority": 1,
    "name": "Transmission",
    "fields": [
        {
            "order": 0,
            "name": "host",
            "label": "Host",
            "value": "transmission",
            "type": "textbox",
            "advanced": false
        },
        {
            "order": 1,
            "name": "port",
            "label": "Port",
            "value": 9091,
            "type": "textbox",
            "advanced": false
        },
        {
            "order": 2,
            "name": "useSsl",
            "label": "Use SSL",
            "helpText": "Use secure connection when connecting to Transmission",
            "value": false,
            "type": "checkbox",
            "advanced": false
        }
    ],
    "implementationName": "Transmission",
    "implementation": "Transmission",
    "configContract": "TransmissionSettings",
    "infoLink": "https://wiki.servarr.com/readarr/supported#transmission",
    "tags": [],
    "id": 3
}`

const addDownloadClient = `{"enable":true,"priority":1,"configContract":"TransmissionSettings","implementation":"Transmission","implementationName":"","name":"Transmission","protocol":"torrent","tags":null,"fields":[{"name":"host","value":"transmission"},{"name":"port","value":9091},{"name":"useSSL","value":false}]}`

const updateDownloadClient = `{"enable":true,"priority":1,"id":3,"configContract":"TransmissionSettings","implementation":"Transmission","implementationName":"","name":"Transmission","protocol":"torrent","tags":null,"fields":[{"name":"host","value":"transmission"},{"name":"port","value":9091},{"name":"useSSL","value":false}]}`

func TestGetDownloadClients(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "downloadClient"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + downloadClientResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*readarr.DownloadClientOutput{
				{
					Enable:             true,
					Priority:           1,
					ID:                 3,
					ConfigContract:     "TransmissionSettings",
					Implementation:     "Transmission",
					ImplementationName: "Transmission",
					InfoLink:           "https://wiki.servarr.com/readarr/supported#transmission",
					Name:               "Transmission",
					Protocol:           "torrent",
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "host",
							Label:    "Host",
							Value:    "transmission",
							Type:     "textbox",
							Advanced: false,
						},
						{
							Order:    1,
							Name:     "port",
							Label:    "Port",
							Value:    float64(9091),
							Type:     "textbox",
							Advanced: false,
						},
						{
							Order:    2,
							Name:     "useSsl",
							Label:    "Use SSL",
							HelpText: "Use secure connection when connecting to Transmission",
							Value:    false,
							Type:     "checkbox",
							Advanced: false,
						},
					},
					Tags: []int{},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*readarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDownloadClients()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetDownloadClient(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "downloadClient", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    downloadClientResponseBody,
			WithRequest:     nil,
			WithResponse: &readarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/readarr/supported#transmission",
				Name:               "Transmission",
				Protocol:           "torrent",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "host",
						Label:    "Host",
						Value:    "transmission",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "port",
						Label:    "Port",
						Value:    float64(9091),
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    2,
						Name:     "useSsl",
						Label:    "Use SSL",
						HelpText: "Use secure connection when connecting to Transmission",
						Value:    false,
						Type:     "checkbox",
						Advanced: false,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*readarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDownloadClient(1)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddDownloadClient(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &readarr.DownloadClientInput{
				Enable:         true,
				Priority:       1,
				ConfigContract: "TransmissionSettings",
				Implementation: "Transmission",
				Name:           "Transmission",
				Protocol:       "torrent",
				Fields: []*starr.FieldInput{
					{
						Name:  "host",
						Value: "transmission",
					},
					{
						Name:  "port",
						Value: 9091,
					},
					{
						Name:  "useSSL",
						Value: false,
					},
				},
			},
			ExpectedRequest: addDownloadClient + "\n",
			ResponseBody:    downloadClientResponseBody,
			WithResponse: &readarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/readarr/supported#transmission",
				Name:               "Transmission",
				Protocol:           "torrent",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "host",
						Label:    "Host",
						Value:    "transmission",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "port",
						Label:    "Port",
						Value:    float64(9091),
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    2,
						Name:     "useSsl",
						Label:    "Use SSL",
						HelpText: "Use secure connection when connecting to Transmission",
						Value:    false,
						Type:     "checkbox",
						Advanced: false,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &readarr.DownloadClientInput{
				Enable:         true,
				Priority:       1,
				ConfigContract: "TransmissionSettings",
				Implementation: "Transmission",
				Name:           "Transmission",
				Protocol:       "torrent",
				Fields: []*starr.FieldInput{
					{
						Name:  "host",
						Value: "transmission",
					},
					{
						Name:  "port",
						Value: 9091,
					},
					{
						Name:  "useSSL",
						Value: false,
					},
				},
			},
			ExpectedRequest: addDownloadClient + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*readarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddDownloadClient(test.WithRequest.(*readarr.DownloadClientInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateDownloadClient(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient", "3?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &readarr.DownloadClientInput{
				Enable:         true,
				Priority:       1,
				ConfigContract: "TransmissionSettings",
				Implementation: "Transmission",
				Name:           "Transmission",
				Protocol:       "torrent",
				Fields: []*starr.FieldInput{
					{
						Name:  "host",
						Value: "transmission",
					},
					{
						Name:  "port",
						Value: 9091,
					},
					{
						Name:  "useSSL",
						Value: false,
					},
				},
				ID: 3,
			},
			ExpectedRequest: updateDownloadClient + "\n",
			ResponseBody:    downloadClientResponseBody,
			WithResponse: &readarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/readarr/supported#transmission",
				Name:               "Transmission",
				Protocol:           "torrent",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "host",
						Label:    "Host",
						Value:    "transmission",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "port",
						Label:    "Port",
						Value:    float64(9091),
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    2,
						Name:     "useSsl",
						Label:    "Use SSL",
						HelpText: "Use secure connection when connecting to Transmission",
						Value:    false,
						Type:     "checkbox",
						Advanced: false,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient", "3?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &readarr.DownloadClientInput{
				Enable:         true,
				Priority:       1,
				ConfigContract: "TransmissionSettings",
				Implementation: "Transmission",
				Name:           "Transmission",
				Protocol:       "torrent",
				Fields: []*starr.FieldInput{
					{
						Name:  "host",
						Value: "transmission",
					},
					{
						Name:  "port",
						Value: 9091,
					},
					{
						Name:  "useSSL",
						Value: false,
					},
				},
				ID: 3,
			},
			ExpectedRequest: updateDownloadClient + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*readarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateDownloadClient(test.WithRequest.(*readarr.DownloadClientInput), false)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteDownloadClient(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "downloadClient", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteDownloadClient(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
