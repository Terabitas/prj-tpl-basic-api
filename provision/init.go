package provision

import (
	"fmt"
	"os"

	"github.com/nildev/lib/registry"
	"gopkg.in/mgo.v2"
)

var (
	TABLE_NAME = "{{.TableName}}"
)

// NildevInitMongoDB init
func NildevInitMongoDB() {
	// In which environment we are
	env := os.Getenv("ND_{{.ApiNameAllCapital}}_ENV")
	fmt.Printf("%s", env)

	session, err := registry.GetMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).Create(&mgo.CollectionInfo{})

	if err != nil {
		fmt.Printf("%s", err)
	}

	usernameIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).EnsureIndex(usernameIndex)

	if err != nil {
		fmt.Printf("%s", err)
	}

	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).EnsureIndex(emailIndex)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func DestroyMongoDB() {
	session, err := registry.GetMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_NAME).DropCollection()

	if err != nil {
		fmt.Printf("%s", err)
	}
}
