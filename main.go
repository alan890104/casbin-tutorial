package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func NewCasbinEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	// clear policy
	e.ClearPolicy()

	// add static roles for each route
	e.AddPermissionForUser("cam-god(site1)", "site1", "/cloud-cameras/*", "*")
	e.AddPermissionForUser("cam-creator(site1)", "site1", "/cloud-cameras/cameras", "POST")
	e.AddPermissionForUser("cam-viewer(site1)", "site1", "/cloud-cameras/cameras/:imei", "GET")
	e.AddPermissionForUser("cam-viewer(site1)", "site1", "/cloud-cameras/cameras/:imei/overview", "GET")
	// define roles that inherit from static roles
	e.AddRoleForUserInDomain("customer(site1)", "cam-creator(site1)", "site1")
	e.AddRoleForUserInDomain("customer(site1)", "cam-viewer(site1)", "site1")
	e.AddRoleForUserInDomain("guard(site1)", "cam-viewer(site1)", "site1")
	e.AddRoleForUserInDomain("god(site1)", "cam-god(site1)", "site1")
	// assign user to roles
	e.AddRoleForUser("alice", "customer(site1)")
	e.AddRoleForUser("bob", "guard(site1)")
	e.AddRoleForUser("god", "god(site1)")
	if err := e.SavePolicy(); err != nil {
		log.Fatal(err)
	}
	return e
}
