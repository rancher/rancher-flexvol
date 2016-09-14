package responsetypes

import (
	"encoding/json"
	"fmt"
)

type DriverOutput struct {
	Status  string
	Message string
	Device  string
}

func (do *DriverOutput) Print() {
	b, _ := json.Marshal(do)
	fmt.Printf("\"%s\"\n", string(b))
}
