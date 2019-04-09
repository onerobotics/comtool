// Package comtool provides the ability to set FANUC robot comments
// via an HTTP request to the KAREL ComSet utility.
package comtool

import (
	"fmt"
	"net/http"
	"net/url"
)

// functionCode is a type for the constants listed below which
// should be provided to Set. The constants correspond to
// the sFc parameter required by ComSet.
type functionCode int

const (
	NUMREG functionCode = 1
	// 2 -- set numreg
	POSREG functionCode = 3
	UALM functionCode = 4
	// 5 -- set ualm_sev
	RIN functionCode = 6
	ROUT functionCode = 7
	DIN functionCode = 8
	DOUT functionCode = 9
	GIN functionCode = 10
	GOUT functionCode = 11
	AIN functionCode = 12
	AOUT functionCode = 13
	SREG functionCode = 14
	// 15 -- set SREG
	// 16 ??
	// 17 ??
	// 18 ??
	FLAG functionCode = 19
)

// Set sets the comment for an item at the provided host.
func Set(code functionCode, id int, comment string, host string) error {
	url := fmt.Sprintf("http://%s/karel/ComSet?sComment=%s&sIndx=%d&sFc=%d", host, url.QueryEscape(comment), id, code)

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
