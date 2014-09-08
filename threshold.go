package nagiosplugin

// Threshold is a combination of a Warning Range and
// a critical Range. One or both of which may be nil.
type Threshold struct {
  Warning   *Range
  Critical  *Range
}

// Returns a new Threshold when given two strings, one for
// the warning range and one for the critical range.
func ParseThresholds( warning string, critical string) (*Threshold, error) {
  warn, err := ParseRange(warning)
  if err != nil {
    warn = nil
  }
  crit, err := ParseRange(critical)
  if err != nil {
    crit = nil
  }
  return &Threshold{
    Warning: warn,
    Critical: crit,
  }, nil
}

// Returns the numerical severity of the value given.
// When no ranges for warning and critical are
// defined it returns UNKNOWN.
func (t *Threshold) Check( value float64) Status {
  if t.Critical == nil && t.Warning == nil {
    return UNKNOWN
  } else if t.Critical != nil && t.Critical.Check( value ) {
    return CRITICAL
  } else if t.Warning != nil && t.Warning.Check( value ) {
    return WARNING
  }
  return OK
}

// Returns true iff the value raises a warning.
func (t *Threshold) IsWarning( value float64) bool {
  if t.Warning.Check( value ) {
    return true
  }
  return false
}

// Returns true iff the value raises a critical.
func (t* Threshold) IsCritical( value float64) bool {
  if t.Critical.Check( value ) {
    return true
  }
  return false
}

