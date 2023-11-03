package customIO

import (
	"io/ioutil"
	"log"

	. "github.com/michalszmidt/cie/objects"
	"gopkg.in/yaml.v3"
)

func PathToYaml(path string) CIE {
	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var cie CIE

	err = yaml.Unmarshal(yfile, &cie)

	if err != nil {
		log.Fatal(err)
	}

	return cie
}
