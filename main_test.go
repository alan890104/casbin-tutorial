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
		// test cam admin
		{"admin", "site1", "/cloud-cameras/cameras", "POST", true},
		{"admin", "site1", "/cloud-cameras/cameras/:imei/overview", "GET", true},
		{"admin", "site1", "/cloud-cameras/cameras/:imei", "DELETE", true},
		{"admin", "site2", "/cloud-cameras/cameras", "POST", false},
		{"admin", "site1", "/cloud-cameras/cameras/:imei", "PUT", true},
		// test distributor
		{"distributor", "site1", "/cloud-cameras/cameras", "POST", true},
		{"distributor", "site1", "/cloud-cameras/cameras/:imei/overview", "GET", true},
		{"distributor", "site1", "/cloud-cameras/cameras/:imei", "DELETE", true},
		{"distributor", "site2", "/cloud-cameras/cameras", "POST", false},
		{"distributor", "site1", "/cloud-cameras/cameras/:imei", "PUT", true},
		{"distributor", "site2", "/devices", "POST", true},
	}

	for _, tc := range testcases {
		actual, err := e.Enforce(tc.user, tc.domain, tc.path, tc.method)
		assert.Nil(t, err)
		assert.Equal(t, tc.expected, actual, "user: %s, path: %s, method: %s", tc.user, tc.path, tc.method)
	}
}
