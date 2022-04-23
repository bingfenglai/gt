package config

import "time"

type CaptchaConfig struct {
	Prefix         string
	ValidityPeriod time.Duration
	NumberCodeLength int
}
