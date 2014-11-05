package nagiosplugin

import (
	"regexp"
)

var boundaryRx = regexp.MustComplie(`/^\^?[+-]?(?:[0-9]+\.?[0-9]*|inf)$/`)

// see https://www.monitoring-plugins.org/doc/new-threshold-syntax.html

// Threshold is a combination of a Warning Range and
// a critical Range. One or both of which may be nil.
type Threshold struct {
	Metrics map[string]Metric
}

type Metric struct {
	Ok       []ThRange
	Warning  []ThRange
	Critical []ThRange
}

type ThRange struct {
	Start          float64
	End            float64
	AlertOnOutside bool
}

// Returns a new Threshold when given two strings, one for
// the warning range and one for the critical range.
func ParseThresholds(thresholds []string) (*Threshold, error) {
	// TODO implement
	metrics := make(map[string]Metric)
	for _, t := range thresholds {
		metricName, metric, err := ParseMetric(t)
		if err != nil {
			return nil, err
		}
		metrics[metricName] = metric
	}

	return &Threshold{
		Metrics: metrics,
	}, nil
}

func ParseMetric(threshold string) (*Metric, error) {
	bounds := strings.Split(threshold, "..")
	if len(bounds) < 2 {
		return nil, errors.New("start and end must be defined")
	}
	// TODO parse whole metric
}

// Returns the numerical severity of the value given.
// When no ranges for warning and critical are
// defined it returns UNKNOWN.
func (t *Threshold) Check(metric string, value float64) Status {
	m, found := t.Metrics[metric]
	if !found {
		// if no levels are speficied return OK
		return OK
	}
	// if an OK level is specified and the value is within range
	// return OK
	if len(m.Ok) > 0 {
		for _, r := range m.Ok {
			if r.Check(value) {
				return OK
			}
		}
	}
	// if an Critical level is specified and the value is within
	// range return Critical
	if len(m.Critical) > 0 {
		for _, r := range m.Critical {
			if r.Check(value) {
				return CRITICAL
			}
		}
	}
	// if an Warning level is specified and the value is within
	// range return Warning
	if len(m.Warning) > 0 {
		for _, r := range m.Warning {
			if r.Check(value) {
				return WARNING
			}
		}
	}
	// if an OK level is specified return CRITICAL
	if len(m.Ok) > 0 {
		return CRITICAL
	}
	return OK
}
