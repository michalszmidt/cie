package objects

// import (
// "gopkg.in/yaml.v3"
// )

// type Settings struct {
// 	Timezone string `yaml:"timezone"`
// }

type CIE struct {
	// Settings Settings `yaml:"settings"`
	ToRemove ToRemove `yaml:"to_remove"`
}

// `yaml:"cie"`

type ToRemove struct {
	Rules []RemoveRule `yaml:"rules"`
}

type RemoveRule struct {
	Regex   string `yaml:"regex"`
	KeyName string `yaml:"key_name"`
}
