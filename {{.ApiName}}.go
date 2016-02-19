package {{.ApiName}} // import "{{.ProjectPath}}"

import (
	"github.com/nildev/lib/registry"
)

// Register method
// @path /custom-register/{provider}
// @method POST
// @query {userName:[a-z]+}
func Register(provider string, userName string) (result bool, err error) {

	// your logic here

	return true, nil
}

// ProtectedResource method which should return only if user passes valid JWT token
// nildev will set `user` variable and will set to it whatever is in JWT Claims token
// @path /protected
// @protected
func ProtectedResource(user registry.User) (result bool, err error) {
	return true, nil
}

