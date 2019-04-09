// Package comtool provides the ability to set FANUC robot comments
// via an HTTP request to the KAREL ComSet utility.
package comtool

import (
	"fmt"
	"net/http"
	"net/url"
)

// FunctionCode is a type for the constants listed below which
// should be provided to Set. The constants correspond to
// the sFc parameter required by ComSet.
type FunctionCode int

const (
	NUMREG FunctionCode = 1
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

// Set sets the comment for an item at the provided host.
func Set(code FunctionCode, id int, comment string, host string) error {
	url := fmt.Sprintf("http://%s/karel/ComSet?sComment=%s&sIndx=%d&sFc=%d", host, url.PathEscape(comment), id, code)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to set comment for code %d id %d comment %s at host %s", code, id, comment, host)
	}

	return nil
}
