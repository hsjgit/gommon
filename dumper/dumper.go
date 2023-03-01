package dumper

import (
	"bytes"
	"encoding/json"
)

func DumpJson(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}

func DumpJsonIndent(obj interface{}) string {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	enc.SetIndent("", "  ")
	enc.Encode(obj)
	return b.String()
}
