package constants

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleEditor  = "editor"
	RoleUser    = "user"
)

var RoleHierarchy = map[string]int{
	RoleUser:    1,
	RoleEditor:  2,
	RoleManager: 3,
	RoleAdmin:   4,
}
