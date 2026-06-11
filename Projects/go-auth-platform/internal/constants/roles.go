package constants

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleEditor  = "editor"
	RoleUser    = "user"
)

var RoleHierarchy = map[string]int{
	RoleUser:    101,
	RoleEditor:  202,
	RoleManager: 303,
	RoleAdmin:   404,
}
