package i18n

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Message map[string]string

func Unmarshal(id, path string) (*Message, error) {
	result := &Message{}
	fileExt := strings.ToLower(filepath.Ext(path))
	if fileExt != ".toml" && fileExt != ".json" && fileExt != ".yaml" {
		return result, fmt.Errorf(fmt.Sprintf(id, "File type not supported"))
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return result, nil
	}

	if strings.HasSuffix(fileExt, ".json") {
		err := json.Unmarshal(buf, result)
		if err != nil {
			return result, err
		}
	}

	if strings.HasSuffix(fileExt, ".yaml") {
		err := yaml.Unmarshal(buf, result)
		if err != nil {
			return result, err
		}
	}

	if strings.HasSuffix(fileExt, ".toml") {
		_, err := toml.Decode(string(buf), result)
		if err != nil {
			return result, err
		}
	}
	return result, nil

}

func Marshal(v interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		err := encoder.Encode(v)
		return buffer.Bytes(), err
	case "toml":
		var buf bytes.Buffer
		enc := toml.NewEncoder(&buf)
		enc.Indent = ""
		err := enc.Encode(v)
		return buf.Bytes(), err
	case "yaml":
		return yaml.Marshal(v)
	}
	return nil, fmt.Errorf("unsupported format: %s", format)
}
