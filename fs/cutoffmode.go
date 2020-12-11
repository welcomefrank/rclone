package fs

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// CutoffMode describes the possible delete modes in the config
type CutoffMode byte

// MaxTransferMode constants
const (
	CutoffModeHard CutoffMode = iota
	CutoffModeSoft
	CutoffModeCautious
	CutoffModeDefault = CutoffModeHard
)

var cutoffModeToString = []string{
	CutoffModeHard:     "HARD",
	CutoffModeSoft:     "SOFT",
	CutoffModeCautious: "CAUTIOUS",
}

// String turns a LogLevel into a string
func (m CutoffMode) String() string {
	if m >= CutoffMode(len(cutoffModeToString)) {
		return fmt.Sprintf("CutoffMode(%d)", m)
	}
	return cutoffModeToString[m]
}

// Set a LogLevel
func (m *CutoffMode) Set(s string) error {
	for n, name := range cutoffModeToString {
		if s != "" && name == strings.ToUpper(s) {
			*m = CutoffMode(n)
			return nil
		}
	}
	return errors.Errorf("Unknown cutoff mode %q", s)
}

// Type of the value
func (m *CutoffMode) Type() string {
	return "string"
}

// UnmarshalJSON makes sure the value can be parsed as a string or integer in JSON
func (x *CutoffMode) UnmarshalJSON(in []byte) error {
	return UnmarshalJSONFlag(in, x, func(i int64) error {
		if i < 0 || i >= int64(len(cutoffModeToString)) {
			return errors.Errorf("Out of range cutoff mode %d", i)
		}
		*x = (CutoffMode)(i)
		return nil
	})
}
