package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/wire"
)

func parseContents(input string) (string, error) {
	r := strings.NewReader(input)
	file, err := wire.NewReader(r).Read()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		inputJSON := args[0].String()

		parsed, err := parseContents(inputJSON)
		if err != nil {
			msg := fmt.Sprintf("unable to parse wire file - %v", err)
			fmt.Print(msg)
			return msg
		}
		pretty, err := prettyJson(parsed)
		if err != nil {
			fmt.Printf("unable to convert wire file to json %s\n", err)
			return "There was an error converting the json"
		}

		return pretty
	})
	return jsonFunc
}

func main() {
	js.Global().Set("parseContents", jsonWrapper())
	<-make(chan bool)
}
