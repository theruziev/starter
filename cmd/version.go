package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/theruziev/starter/internal/pkg/info"
)

type versionCli struct{}

func (v *versionCli) Run(_ *contextCli) error {
	resultJSON, err := json.Marshal(info.Information())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resultJSON))
	return nil
}
