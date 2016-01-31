package {{.ApiName}} // import "{{.ProjectPath}}"

import (
	"{{.ProjectPath}}/provision"
	"github.com/nildev/lib/registry"
)

func Register(userName string, email string) (result bool, err error) {

	if err != nil {
		return false, err
	}

	acc := map[string]string{"u":userName, "e":email}

	session, err := registry.GetMongoDBClient()
	if err != nil {
		return false, err
	}

	collection := session.DB(registry.GetDatabaseName()).C(provision.TABLE_NAME)
	err = collection.Insert(acc)

	if err != nil {
		return false, err
	}

	return true, nil
}

func AuthByEmail(email string, password string) (result bool, err error) {
	return true, nil
}

