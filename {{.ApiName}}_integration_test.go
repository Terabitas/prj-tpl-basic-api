// +build integration

package {{.ApiName}}

import (
	. "gopkg.in/check.v1"
	"{{.ProjectPath}}/provision"
)

type {{.ApiNameExported}}IntegrationSuite struct{}

var _ = Suite(&{{.ApiNameExported}}IntegrationSuite{})

func TestMain(m *testing.M) {
	provision.NildevInitMongoDB()
	code := m.Run()
	provision.DestroyMongoDB()
	os.Exit(code)
}

func (s *{{.ApiNameExported}}IntegrationSuite) TestIfSomethingGoodHappensWithinDB(c *C) {

	rez, err := Register("username", "password")
	c.Assert(rez, Equals, true)
	c.Assert(err, IsNil)
}
