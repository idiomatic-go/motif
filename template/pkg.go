package template

import "reflect"

type pkg struct{}

var pkgPath = reflect.TypeOf(any(pkg{})).PkgPath()
