package radarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/radarr"
	"golift.io/starr/starrtest"
)

const downloadClientResponseBody = `{
    "enable": true,
    "protocol": "torrent",
    "priority": 1,
    "removeCompletedDownloads": false,
    "removeFailedDownloads": false,
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
    "infoLink": "https://wiki.servarr.com/radarr/supported#transmission",
    "tags": [],
    "id": 3
}`

const addDownloadClient = `{"enable":true,"removeCompletedDownloads":false,"removeFailedDownloads":false,` +
	`"priority":1,"configContract":"TransmissionSettings","implementation":"Transmission","name":"Transmission",` +
	`"protocol":"torrent","tags":null,"fields":[{"name":"host","value":"transmission"},` +
	`{"name":"port","value":9091},{"name":"useSSL","value":false}]}`

const updateDownloadClient = `{"enable":true,"removeCompletedDownloads":false,"removeFailedDownloads":false,` +
	`"priority":1,"id":3,"configContract":"TransmissionSettings","implementation":"Transmission","name":"Transmission",` +
	`"protocol":"torrent","tags":null,"fields":[{"name":"host","value":"transmission"},` +
	`{"name":"port","value":9091},{"name":"useSSL","value":false}]}`

func TestGetDownloadClients(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "downloadClient"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + downloadClientResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*radarr.DownloadClientOutput{
				{
					Enable:             true,
					Priority:           1,
					ID:                 3,
					ConfigContract:     "TransmissionSettings",
					Implementation:     "Transmission",
					ImplementationName: "Transmission",
					InfoLink:           "https://wiki.servarr.com/radarr/supported#transmission",
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*radarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "downloadClient", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    downloadClientResponseBody,
			WithRequest:     nil,
			WithResponse: &radarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/radarr/supported#transmission",
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*radarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &radarr.DownloadClientInput{
				Enable:                   true,
				RemoveCompletedDownloads: false,
				RemoveFailedDownloads:    false,
				Priority:                 1,
				ConfigContract:           "TransmissionSettings",
				Implementation:           "Transmission",
				Name:                     "Transmission",
				Protocol:                 "torrent",
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
			WithResponse: &radarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/radarr/supported#transmission",
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &radarr.DownloadClientInput{
				Enable:                   true,
				RemoveCompletedDownloads: false,
				RemoveFailedDownloads:    false,
				Priority:                 1,
				ConfigContract:           "TransmissionSettings",
				Implementation:           "Transmission",
				Name:                     "Transmission",
				Protocol:                 "torrent",
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
			WithResponse:    (*radarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddDownloadClient(test.WithRequest.(*radarr.DownloadClientInput))
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient", "3?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &radarr.DownloadClientInput{
				Enable:                   true,
				RemoveCompletedDownloads: false,
				RemoveFailedDownloads:    false,
				Priority:                 1,
				ConfigContract:           "TransmissionSettings",
				Implementation:           "Transmission",
				Name:                     "Transmission",
				Protocol:                 "torrent",
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
			WithResponse: &radarr.DownloadClientOutput{
				Enable:             true,
				Priority:           1,
				ID:                 3,
				ConfigContract:     "TransmissionSettings",
				Implementation:     "Transmission",
				ImplementationName: "Transmission",
				InfoLink:           "https://wiki.servarr.com/radarr/supported#transmission",
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient", "3?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &radarr.DownloadClientInput{
				Enable:                   true,
				RemoveCompletedDownloads: false,
				RemoveFailedDownloads:    false,
				Priority:                 1,
				ConfigContract:           "TransmissionSettings",
				Implementation:           "Transmission",
				Name:                     "Transmission",
				Protocol:                 "torrent",
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
			WithResponse:    (*radarr.DownloadClientOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateDownloadClient(test.WithRequest.(*radarr.DownloadClientInput), false)
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "downloadClient", "2"),
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
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteDownloadClient(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
