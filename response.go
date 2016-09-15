package flexvol

import (
	"encoding/json"
	"fmt"
)

type DriverOutput struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Device  string `json:"device"`
}

func (do *DriverOutput) Print() {
	b, _ := json.Marshal(do)
	fmt.Printf("\"%s\"\n", string(b))
}
