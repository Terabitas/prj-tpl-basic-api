package {{.ApiName}}

import (
	. "gopkg.in/check.v1"
)

type {{.ApiNameExported}}Suite struct{}

var _ = Suite(&{{.ApiNameExported}}Suite{})

func (s *{{.ApiNameExported}}Suite) TestIfSomethingGoodHappens(c *C) {

}
