package features

import (
	"github.com/DATA-DOG/godog"
)

// Feat it's the expected struct to export your steps
type Feat struct {
	expr interface{}
	step interface{}
}

func FeatureContext(s *godog.Suite) {
	ctx := &Context{}
	s.BeforeScenario(func(interface{}) {
		ctx.m = make(map[string]interface{})
	})

	for _, feat := range CmdFeatures(ctx) {
		s.Step(feat.expr, feat.step)
	}
}
