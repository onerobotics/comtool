// Package comtool provides the ability to set FANUC robot comments
// via an HTTP request to the KAREL ComSet utility.
package comtool

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// FunctionCode is a type for the constants listed below which
// should be provided to Set. The constants correspond to
// the sFc parameter required by ComSet.
type FunctionCode int

const (
	INVALID FunctionCode = 0
	NUMREG  FunctionCode = 1
	// 2 -- set numreg
	POSREG FunctionCode = 3
	UALM   FunctionCode = 4
	// 5 -- set ualm_sev
	RIN  FunctionCode = 6
	ROUT FunctionCode = 7
	DIN  FunctionCode = 8
	DOUT FunctionCode = 9
	GIN  FunctionCode = 10
	GOUT FunctionCode = 11
	AIN  FunctionCode = 12
	AOUT FunctionCode = 13
	SREG FunctionCode = 14
	// 15 -- set SREG
	// 16 ??
	// 17 ??
	// 18 ??
	FLAG FunctionCode = 19
)

var codes = [...]string{
	NUMREG: "R",
	POSREG: "PR",
	UALM:   "UALM",
	RIN:    "RI",
	ROUT:   "RO",
	DIN:    "DI",
	DOUT:   "DO",
	GIN:    "GI",
	GOUT:   "GO",
	AIN:    "AI",
	AOUT:   "AO",
	SREG:   "SR",
	FLAG:   "F",
}

func (f FunctionCode) String() string {
	s := codes[f]
	if s != "" {
		return s
	}

	return "functionCode(" + strconv.Itoa(int(f)) + ")"
}

var (
	ErrForbidden    = errForbidden()
	ErrUnauthorized = errUnauthorized()
)

func errForbidden() error {
	return errors.New("Forbidden. Please unlock KAREL via Setup > Host Comm > HTTP.")
}

func errUnauthorized() error {
	return errors.New("Unauthorized. Please unlock KAREL via Setup > Host Comm > HTTP.")
}

// Set sets the comment for an item at the provided host.
func Set(code FunctionCode, id int, comment string, host string, timeout time.Duration) error {
	url := fmt.Sprintf("http://%s/karel/ComSet?sComment=%s&sIndx=%d&sFc=%d", host, url.PathEscape(comment), id, code)

	client := http.Client{
		Timeout: time.Duration(timeout),
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusForbidden:
			return ErrForbidden
		case http.StatusUnauthorized:
			return ErrUnauthorized
		default:
			return fmt.Errorf("Failed to set comment for %s[%d], '%s', at host %s: %d", code, id, comment, host, resp.StatusCode)
		}
	}

	return nil
}
