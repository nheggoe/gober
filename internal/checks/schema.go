package checks

import (
	"iter"
	"net/url"
	"time"

	"github.com/nheggoe/gober/internal/util"
)

type HealthConfig struct {
	DefaultInterval string        `hcl:"interval,optional"`
	DefaultTimeout  string        `hcl:"timeout,optional"`
	Checks          []CheckConfig `hcl:"check,block"`
}

type CheckConfig struct {
	Name     string `hcl:"name,label"`
	Target   string `hcl:"target"`
	PushURL  string `hcl:"push_url"`
	Interval string `hcl:"interval,optional"`
	Timeout  string `hcl:"timeout,optional"`
}

func (f HealthConfig) ToRuntimeConfig() (_ RuntimeConfig, err error) {
	defer util.WrapError(&err, "ToModel")
	errs := new(util.Errors)

	const (
		fallbackInterval = 2 * time.Hour
		fallbackTimeout  = 8 * time.Second
	)

	defaultInterval, err := parseDurationOr(f.DefaultInterval, fallbackInterval)
	if err != nil {
		errs.Append(err)
	}

	defaultTimeout, err := parseDurationOr(f.DefaultTimeout, fallbackTimeout)
	if err != nil {
		errs.Append(err)
	}

	// cannot proceed without default values
	if !errs.IsEmpty() {
		return RuntimeConfig{}, errs
	}

	checks, err := util.MapError(f.Checks, func(cf CheckConfig) (Check, error) {
		checkErrs := new(util.Errors)
		target, err := url.Parse(cf.Target)
		if err != nil {
			checkErrs.Append(err)
		}
		pushURL, err := url.Parse(cf.PushURL)
		if err != nil {
			checkErrs.Append(err)
		}
		interval, err := parseDurationOr(cf.Interval, defaultInterval)
		if err != nil {
			checkErrs.Append(err)
		}
		timeout, err := parseDurationOr(cf.Timeout, defaultTimeout)
		if err != nil {
			checkErrs.Append(err)
		}
		if !checkErrs.IsEmpty() {
			return Check{}, checkErrs
		}
		return Check{
			Name:     cf.Name,
			Target:   target,
			PushURL:  pushURL,
			Interval: interval,
			Timeout:  timeout,
		}, nil
	})
	if err != nil {
		errs.Append(err)
	}

	if !errs.IsEmpty() {
		return RuntimeConfig{}, errs
	}
	return RuntimeConfig{
		DefaultInterval: defaultInterval,
		DefaultTimeout:  defaultTimeout,
		Checks:          checks,
	}, nil
}

type RuntimeConfig struct {
	DefaultInterval time.Duration
	DefaultTimeout  time.Duration
	Checks          []Check
}
type Check struct {
	Name     string
	Target   *url.URL
	PushURL  *url.URL
	Interval time.Duration
	Timeout  time.Duration
}

type Checks []Check

func (cs *Checks) All() iter.Seq[Check] {
	return func(yield func(Check) bool) {
		for _, c := range *cs {
			if !yield(c) {
				return
			}
		}
	}
}

// ========	helper ========

func parseDurationOr(s string, fallback time.Duration) (time.Duration, error) {
	if s == "" {
		return fallback, nil
	}

	parsed, err := time.ParseDuration(s)
	if err != nil {
		var zero time.Duration
		return zero, err
	}

	return parsed, nil
}
