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

const (
	firstRootFolder = `{
		"path": "/movies",
		"accessible": true,
		"freeSpace": 252221177856,
		"unmappedFolders": [],
		"id": 1
	}`
	secondRootFolder = `{
		"path": "/collections",
		"accessible": true,
		"freeSpace": 252221177856,
		"unmappedFolders": [
			{
				"name": "1",
				"path": "/collections/1"
			}
		],
		"id": 2
	}`
)

func TestGetRootFolders(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[` + firstRootFolder + `,` + secondRootFolder + `]`,
			WithResponse: []*radarr.RootFolder{
				{
					Path:            "/movies",
					Accessible:      true,
					FreeSpace:       252221177856,
					UnmappedFolders: []*starr.Path{},
					ID:              1,
				},
				{
					Path:       "/collections",
					Accessible: true,
					FreeSpace:  252221177856,
					UnmappedFolders: []*starr.Path{
						{
							Name: "1",
							Path: "/collections/1",
						},
					},
					ID: 2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   []*radarr.RootFolder(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRootFolders()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetRootFolder(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    int64(1),
			ResponseBody:   firstRootFolder,
			WithResponse: &radarr.RootFolder{
				Path:            "/movies",
				Accessible:      true,
				FreeSpace:       252221177856,
				UnmappedFolders: []*starr.Path{},
				ID:              1,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    int64(1),
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   (*radarr.RootFolder)(nil),
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRootFolder(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddRootFolder(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "201",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder"),
			ExpectedMethod: "POST",
			ResponseStatus: 201,
			WithRequest: &radarr.RootFolder{
				Path: "/collections",
			},
			ExpectedRequest: `{"path":"/collections"}` + "\n",
			ResponseBody:    secondRootFolder,
			WithResponse: &radarr.RootFolder{
				Path:       "/collections",
				Accessible: true,
				FreeSpace:  252221177856,
				UnmappedFolders: []*starr.Path{
					{
						Name: "1",
						Path: "/collections/1",
					},
				},
				ID: 2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &radarr.RootFolder{
				Path: "/collections",
			},
			ExpectedRequest: `{"path":"/collections"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*radarr.RootFolder)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddRootFolder(test.WithRequest.(*radarr.RootFolder))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteRootFolder(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "rootFolder", "2"),
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
			err := client.DeleteRootFolder(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
