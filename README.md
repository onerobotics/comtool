# Deprecated. Use [go-fanuc](https://github.com/onerobotics/go-fanuc) instead.

# comtool

Set fanuc data comments via the controller KAREL ComSet tool.

## Usage

    package main

    import (
    	"fmt"

    	"github.com/onerobotics/comtool"
    )

    func main() {
    	err := comtool.Set(comtool.NUMREG, 1, "test", "127.0.0.101")
    	if err != nil {
    		fmt.Println(err)
    	}
    }
