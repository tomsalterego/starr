package prowlarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/starrtest"
)

//nolint:lll,nolintlint // go linters are pretty stupid sometimes.
const (
	notificationResponseBody = `{
    "onHealthIssue": false,
    "onApplicationUpdate": true,
    "supportsOnHealthIssue": true,
    "includeHealthWarnings": false,
    "supportsOnApplicationUpdate": true,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/prowlarr.sh",
		"type": "filePath",
		"advanced": false
	  },
	  {
		"order": 1,
		"name": "arguments",
		"label": "Arguments",
		"helpText": "Arguments to pass to the script",
		"type": "textbox",
		"advanced": false,
		"hidden": "hiddenIfNotSet"
	  }
	],
	"implementationName": "Custom Script",
	"implementation": "CustomScript",
	"configContract": "CustomScriptSettings",
	"infoLink": "https://wiki.servarr.com/prowlarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`
	addNotification    = `{"onGrab":false,"onHealthIssue":false,"onHealthRestored":false,"onApplicationUpdate":true,"supportsOnGrab":false,"includeManualGrabs":false,"supportsOnHealthIssue":false,"supportsOnHealthRestored":false,"includeHealthWarnings":false,"supportsOnApplicationUpdate":false,"name":"Test","implementationName":"","implementation":"CustomScript","configContract":"CustomScriptSettings","infoLink":"","tags":null,"fields":[{"name":"path","value":"/scripts/prowlarr.sh"}]}`
	updateNotification = `{"onGrab":false,"onHealthIssue":false,"onHealthRestored":false,"onApplicationUpdate":true,"supportsOnGrab":false,"includeManualGrabs":false,"supportsOnHealthIssue":false,"supportsOnHealthRestored":false,"includeHealthWarnings":false,"supportsOnApplicationUpdate":false,"id":3,"name":"Test","implementationName":"","implementation":"CustomScript","configContract":"CustomScriptSettings","infoLink":"","tags":null,"fields":[{"name":"path","value":"/scripts/prowlarr.sh"}]}`
)

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*prowlarr.NotificationOutput{
				{
					OnApplicationUpdate:         true,
					SupportsOnHealthIssue:       true,
					SupportsOnApplicationUpdate: true,
					ID:                          3,
					Name:                        "Test",
					ImplementationName:          "Custom Script",
					Implementation:              "CustomScript",
					ConfigContract:              "CustomScriptSettings",
					InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
					Tags:                        []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/prowlarr.sh",
							Type:     "filePath",
							Advanced: false,
						},
						{
							Order:    1,
							Name:     "arguments",
							Label:    "Arguments",
							HelpText: "Arguments to pass to the script",
							Hidden:   "hiddenIfNotSet",
							Type:     "textbox",
							Advanced: false,
						},
					},
					Message: struct {
						Message string `json:"message"`
						Type    string `json:"type"`
					}{
						Message: "Testing will execute the script with the EventType set to Test",
						Type:    "warning",
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotifications()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
				Message: struct {
					Message string `json:"message"`
					Type    string `json:"type"`
				}{
					Message: "Testing will execute the script with the EventType set to Test",
					Type:    "warning",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotification(1)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
				Message: struct {
					Message string `json:"message"`
					Type    string `json:"type"`
				}{
					Message: "Testing will execute the script with the EventType set to Test",
					Type:    "warning",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*prowlarr.NotificationInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				ID:                  3,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
				Message: struct {
					Message string `json:"message"`
					Type    string `json:"type"`
				}{
					Message: "Testing will execute the script with the EventType set to Test",
					Type:    "warning",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				ID:                  3,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*prowlarr.NotificationInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "2"),
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
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
