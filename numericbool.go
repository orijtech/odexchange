package odexchange

// NumericBool represents an interoperable bool <-- int

import (
	"encoding/json"
	"strconv"
)

type NumericBool bool

func (bb *NumericBool) UnmarshalJSON(blob []byte) error {
	if len(blob) < 1 {
		*bb = false
		return nil
	}

	s := string(blob)
	// Try first parsing an integer
	if pBool, err := strconv.ParseBool(s); err == nil {
		*bb = NumericBool(pBool)
	} else if pInt, err := strconv.ParseInt(s, 10, 32); err == nil {
		*bb = pInt != 0
	}

	return nil
}

func (bb NumericBool) MarshalJSON() ([]byte, error) {
	v := 0
	if bb {
		v = 1
	}
	return json.Marshal(v)
}
