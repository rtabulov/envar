package envar

import (
	"io/ioutil"
	"os"
	"strings"
)

// Load func
func Load() error {
	bts, err := ioutil.ReadFile("./.env")
	if err != nil {
		return err
	}

	str := string(bts)
	fields := strings.Fields(str)

	for _, f := range fields {
		splt := strings.SplitN(f, "=", 2)
		key, value := splt[0], splt[1]
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
