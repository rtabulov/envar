package autoload

import "github.com/rtabulov/envar"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	must(envar.Load())
}
