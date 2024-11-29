package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 整形したJSONを標準出力に出力する
func PrintJson(v interface{}) {
	in, err := json.Marshal(v)
	if err == nil {
		var out bytes.Buffer
		err = json.Indent(&out, in, "", "  ")
		if err == nil {
			fmt.Println(out.String())
		}
	} else {
		fmt.Println(err)
	}
}
