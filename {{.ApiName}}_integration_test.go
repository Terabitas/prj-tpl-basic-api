// +build integration

package {{.ApiName}}

import (
	"os"
	"testing"
	. "gopkg.in/check.v1"
	"{{.ProjectPath}}/provision"
)

type {{.ApiNameExported}}IntegrationSuite struct{}

var _ = Suite(&{{.ApiNameExported}}IntegrationSuite{})

func TestMain(m *testing.M) {
	dbname := "nildev"
	ip := os.Getenv("ND_IP_PORT")

	os.Setenv("ND_ENV", "dev")
	os.Setenv("ND_DATABASE_NAME", dbname)
	os.Setenv("ND_MONGODB_URL", "mongodb://"+ip+"/"+dbname)

	provision.NildevInitMongoDB()
	code := m.Run()
	provision.DestroyMongoDB()
	os.Exit(code)
}

func (s *{{.ApiNameExported}}IntegrationSuite) TestIfSomethingGoodHappensWithinDB(c *C) {
	rez, err := Register("github", "nildevuser")
	c.Assert(rez, Equals, true)
	c.Assert(err, IsNil)
}
