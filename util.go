package vonage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"path"
)

func mergeStructsAsJson(structs ...interface{}) io.Reader {
	var m map[string]string
	for _, s := range structs {
		b, _ := json.Marshal(s)
		json.Unmarshal(b, &m)
	}
	b, _ := json.Marshal(m)
	return bytes.NewReader(b)
}

// util function to join url.
// Without this function, scheme(ex. https) would lack or unexpected slash would appear.
func uriJoin(parent, child string) (string, error) {
	u, err := url.Parse(parent)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, child)
	return fmt.Sprint(u), nil
}
