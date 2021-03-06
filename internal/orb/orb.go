package orb

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Orb struct {
	URL     string
	Repokey string
}

func ParseOrbConfig(orbConfig string) (*Orb, error) {
	gitOrbConfig, err := ioutil.ReadFile(orbConfig)

	if err != nil {
		return nil, errors.Wrap(err, "unable to read orbconfig")
	}

	orb := &Orb{}
	if err := yaml.Unmarshal(gitOrbConfig, orb); err != nil {
		return nil, errors.Wrap(err, "unable to parse orbconfig")
	}

	return orb, nil
}
