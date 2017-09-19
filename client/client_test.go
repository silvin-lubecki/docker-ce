package client // import "github.com/docker/docker/client"

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"testing"

	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/internal/testutil"
	"github.com/gotestyourself/gotestyourself/env"
	"github.com/gotestyourself/gotestyourself/skip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEnvClient(t *testing.T) {
	skip.If(t, runtime.GOOS == "windows")

	testcases := []struct {
		doc             string
		envs            map[string]string
		expectedError   string
		expectedVersion string
	}{
		{
			doc:             "default api version",
			envs:            map[string]string{},
			expectedVersion: api.DefaultVersion,
		},
		{
			doc: "invalid cert path",
			envs: map[string]string{
				"DOCKER_CERT_PATH": "invalid/path",
			},
			expectedError: "Could not load X509 key pair: open invalid/path/cert.pem: no such file or directory",
		},
		{
			doc: "default api version with cert path",
			envs: map[string]string{
				"DOCKER_CERT_PATH": "testdata/",
			},
			expectedVersion: api.DefaultVersion,
		},
		{
			doc: "default api version with cert path and tls verify",
			envs: map[string]string{
				"DOCKER_CERT_PATH":  "testdata/",
				"DOCKER_TLS_VERIFY": "1",
			},
			expectedVersion: api.DefaultVersion,
		},
		{
			doc: "default api version with cert path and host",
			envs: map[string]string{
				"DOCKER_CERT_PATH": "testdata/",
				"DOCKER_HOST":      "https://notaunixsocket",
			},
			expectedVersion: api.DefaultVersion,
		},
		{
			doc: "invalid docker host",
			envs: map[string]string{
				"DOCKER_HOST": "host",
			},
			expectedError: "unable to parse docker host `host`",
		},
		{
			doc: "invalid docker host, with good format",
			envs: map[string]string{
				"DOCKER_HOST": "invalid://url",
			},
			expectedVersion: api.DefaultVersion,
		},
		{
			doc: "override api version",
			envs: map[string]string{
				"DOCKER_API_VERSION": "1.22",
			},
			expectedVersion: "1.22",
		},
	}

	defer env.PatchAll(t, nil)()
	for _, c := range testcases {
		env.PatchAll(t, c.envs)
		apiclient, err := NewEnvClient()
		if c.expectedError != "" {
			assert.Error(t, err, c.doc)
			assert.Equal(t, c.expectedError, err.Error(), c.doc)
		} else {
			assert.NoError(t, err, c.doc)
			version := apiclient.ClientVersion()
			assert.Equal(t, c.expectedVersion, version, c.doc)
		}

		if c.envs["DOCKER_TLS_VERIFY"] != "" {
			// pedantic checking that this is handled correctly
			tr := apiclient.client.Transport.(*http.Transport)
			assert.NotNil(t, tr.TLSClientConfig, c.doc)
			assert.Equal(t, tr.TLSClientConfig.InsecureSkipVerify, false, c.doc)
		}
	}
}

func TestGetAPIPath(t *testing.T) {
	testcases := []struct {
		version  string
		path     string
		query    url.Values
		expected string
	}{
		{"", "/containers/json", nil, "/containers/json"},
		{"", "/containers/json", url.Values{}, "/containers/json"},
		{"", "/containers/json", url.Values{"s": []string{"c"}}, "/containers/json?s=c"},
		{"1.22", "/containers/json", nil, "/v1.22/containers/json"},
		{"1.22", "/containers/json", url.Values{}, "/v1.22/containers/json"},
		{"1.22", "/containers/json", url.Values{"s": []string{"c"}}, "/v1.22/containers/json?s=c"},
		{"v1.22", "/containers/json", nil, "/v1.22/containers/json"},
		{"v1.22", "/containers/json", url.Values{}, "/v1.22/containers/json"},
		{"v1.22", "/containers/json", url.Values{"s": []string{"c"}}, "/v1.22/containers/json?s=c"},
		{"v1.22", "/networks/kiwl$%^", nil, "/v1.22/networks/kiwl$%25%5E"},
	}

	for _, testcase := range testcases {
		c := Client{version: testcase.version, basePath: "/"}
		actual := c.getAPIPath(testcase.path, testcase.query)
		assert.Equal(t, actual, testcase.expected)
	}
}

func TestParseHostURL(t *testing.T) {
	testcases := []struct {
		host        string
		expected    *url.URL
		expectedErr string
	}{
		{
			host:        "",
			expectedErr: "unable to parse docker host",
		},
		{
			host:        "foobar",
			expectedErr: "unable to parse docker host",
		},
		{
			host:     "foo://bar",
			expected: &url.URL{Scheme: "foo", Host: "bar"},
		},
		{
			host:     "tcp://localhost:2476",
			expected: &url.URL{Scheme: "tcp", Host: "localhost:2476"},
		},
		{
			host:     "tcp://localhost:2476/path",
			expected: &url.URL{Scheme: "tcp", Host: "localhost:2476", Path: "/path"},
		},
	}

	for _, testcase := range testcases {
		actual, err := ParseHostURL(testcase.host)
		if testcase.expectedErr != "" {
			testutil.ErrorContains(t, err, testcase.expectedErr)
		}
		assert.Equal(t, testcase.expected, actual)
	}
}

func TestNewEnvClientSetsDefaultVersion(t *testing.T) {
	defer env.PatchAll(t, map[string]string{
		"DOCKER_HOST":        "",
		"DOCKER_API_VERSION": "",
		"DOCKER_TLS_VERIFY":  "",
		"DOCKER_CERT_PATH":   "",
	})()

	client, err := NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, client.version, api.DefaultVersion)

	expected := "1.22"
	os.Setenv("DOCKER_API_VERSION", expected)
	client, err = NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, client.version)
}

// TestNegotiateAPIVersionEmpty asserts that client.Client can
// negotiate a compatible APIVersion when omitted
func TestNegotiateAPIVersionEmpty(t *testing.T) {
	defer env.PatchAll(t, map[string]string{"DOCKER_API_VERSION": ""})

	client, err := NewEnvClient()
	require.NoError(t, err)

	ping := types.Ping{
		APIVersion:   "",
		OSType:       "linux",
		Experimental: false,
	}

	// set our version to something new
	client.version = "1.25"

	// if no version from server, expect the earliest
	// version before APIVersion was implemented
	expected := "1.24"

	// test downgrade
	client.NegotiateAPIVersionPing(ping)
	assert.Equal(t, expected, client.version)
}

// TestNegotiateAPIVersion asserts that client.Client can
// negotiate a compatible APIVersion with the server
func TestNegotiateAPIVersion(t *testing.T) {
	client, err := NewEnvClient()
	require.NoError(t, err)

	expected := "1.21"
	ping := types.Ping{
		APIVersion:   expected,
		OSType:       "linux",
		Experimental: false,
	}

	// set our version to something new
	client.version = "1.22"

	// test downgrade
	client.NegotiateAPIVersionPing(ping)
	assert.Equal(t, expected, client.version)

	// set the client version to something older, and verify that we keep the
	// original setting.
	expected = "1.20"
	client.version = expected
	client.NegotiateAPIVersionPing(ping)
	assert.Equal(t, expected, client.version)

}

// TestNegotiateAPIVersionOverride asserts that we honor
// the environment variable DOCKER_API_VERSION when negotianing versions
func TestNegotiateAPVersionOverride(t *testing.T) {
	expected := "9.99"
	defer env.PatchAll(t, map[string]string{"DOCKER_API_VERSION": expected})()

	client, err := NewEnvClient()
	require.NoError(t, err)

	ping := types.Ping{
		APIVersion:   "1.24",
		OSType:       "linux",
		Experimental: false,
	}

	// test that we honored the env var
	client.NegotiateAPIVersionPing(ping)
	assert.Equal(t, expected, client.version)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (rtf roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rtf(req)
}

type bytesBufferClose struct {
	*bytes.Buffer
}

func (bbc bytesBufferClose) Close() error {
	return nil
}

func TestClientRedirect(t *testing.T) {
	client := &http.Client{
		CheckRedirect: CheckRedirect,
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.URL.String() == "/bla" {
				return &http.Response{StatusCode: 404}, nil
			}
			return &http.Response{
				StatusCode: 301,
				Header:     map[string][]string{"Location": {"/bla"}},
				Body:       bytesBufferClose{bytes.NewBuffer(nil)},
			}, nil
		}),
	}

	cases := []struct {
		httpMethod  string
		expectedErr error
		statusCode  int
	}{
		{http.MethodGet, nil, 301},
		{http.MethodPost, &url.Error{Op: "Post", URL: "/bla", Err: ErrRedirect}, 301},
		{http.MethodPut, &url.Error{Op: "Put", URL: "/bla", Err: ErrRedirect}, 301},
		{http.MethodDelete, &url.Error{Op: "Delete", URL: "/bla", Err: ErrRedirect}, 301},
	}

	for _, tc := range cases {
		req, err := http.NewRequest(tc.httpMethod, "/redirectme", nil)
		assert.NoError(t, err)
		resp, err := client.Do(req)
		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.statusCode, resp.StatusCode)
	}
}
