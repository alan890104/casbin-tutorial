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
	e.AddPermissionForUser("cam-admin(site1)", "site1", "/cloud-cameras/*", "*")
	e.AddPermissionForUser("cam-creator(site1)", "site1", "/cloud-cameras/cameras", "POST")
	e.AddPermissionForUser("cam-viewer(site1)", "site1", "/cloud-cameras/cameras/:imei", "GET")
	e.AddPermissionForUser("cam-viewer(site1)", "site1", "/cloud-cameras/cameras/:imei/overview", "GET")
	e.AddPermissionForUser("device-admin(site2)", "site2", "/devices*", "*")
	// define roles that inherit from static roles
	e.AddRoleForUserInDomain("customer(site1)", "cam-creator(site1)", "site1")
	e.AddRoleForUserInDomain("customer(site1)", "cam-viewer(site1)", "site1")
	e.AddRoleForUserInDomain("guard(site1)", "cam-viewer(site1)", "site1")
	e.AddRoleForUserInDomain("admin(site1)", "cam-admin(site1)", "site1")
	e.AddRoleForUserInDomain("admin(site2)", "device-admin(site2)", "site2")
	// assign user to roles
	e.AddRoleForUser("alice", "customer(site1)")
	e.AddRoleForUser("bob", "guard(site1)")
	e.AddRoleForUser("admin", "admin(site1)")
	e.AddRoleForUser("admin", "admin(site2)")
	// add business entity roles
	e.AddRoleForUser("distributor", "admin")
	// add headquarter roles
	e.AddRoleForUser("headquarter", "distributor")

	if err := e.SavePolicy(); err != nil {
		log.Fatal(err)
	}
	return e
}
