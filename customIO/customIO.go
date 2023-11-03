package customIO

import (
    . "github.com/michalszmidt/cie/objects"
    "gopkg.in/yaml.v3"
        "io/ioutil"
    "os"
    )

func  PathToYaml(path string) CIE {
    var cie CIE     
    file, _ := os.Open(path)
    defer file.Close()
    data, _ := ioutil.ReadAll(file)
    yaml.Unmarshal(data, cie)
    return cie
}