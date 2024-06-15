package configloader

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func Load(configPath string) *koanf.Koanf {
	k := koanf.New(".")

	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)

	}

	return k
}
