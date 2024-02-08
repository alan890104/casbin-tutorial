package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicy(t *testing.T) {
	e := NewCasbinEnforcer()

	testcases := []struct {
		user     string
		domain   string
		path     string
		method   string
		expected bool
	}{
		{"alice", "site1", "/cloud-cameras/cameras", "POST", true},
		{"alice", "site1", "/cloud-cameras/cameras/P02W00484", "GET", true},
		{"alice", "site1", "/cloud-cameras/cameras/P02W00484/overview", "GET", true},
		{"alice", "site1", "/cloud-cameras/cameras/P02W00484/overview/kkk", "GET", false},
		{"bob", "site1", "/cloud-cameras/cameras/P02W00484", "GET", true},
		{"bob", "site1", "/cloud-cameras/cameras", "POST", false},
		{"alice", "site2", "/cloud-cameras/cameras", "POST", false},
		{"alice", "site2", "/cloud-cameras/cameras", "DELETE", false},
		// test cam god
		{"god", "site1", "/cloud-cameras/cameras", "POST", true},
		{"god", "site1", "/cloud-cameras/cameras/:imei/overview", "GET", true},
		{"god", "site1", "/cloud-cameras/cameras/:imei", "DELETE", true},
		{"god", "site2", "/cloud-cameras/cameras", "POST", false},
		{"god", "site1", "/cloud-cameras/cameras/:imei", "PUT", true},
	}

	for _, tc := range testcases {
		actual, err := e.Enforce(tc.user, tc.domain, tc.path, tc.method)
		assert.Nil(t, err)
		assert.Equal(t, tc.expected, actual, "user: %s, path: %s, method: %s", tc.user, tc.path, tc.method)
	}
}
