package provision

import (
	"fmt"

	"github.com/nildev/lib/registry"
	"gopkg.in/mgo.v2"
)

var (
	TABLE_NAME = "{{.TableName}}"
)

// NildevInitMongoDB init
func NildevInitMongoDB() {
	// In which environment we are
	env := registry.GetEnv()
	fmt.Printf("%s", env)

	session, err := registry.CreateMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).Create(&mgo.CollectionInfo{})

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func DestroyMongoDB() {
	session, err := registry.CreateMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).DropCollection()

	if err != nil {
		fmt.Printf("%s", err)
	}
}
