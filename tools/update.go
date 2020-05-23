package i18n

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Rain-31/go-i18n/v1/i18n"
)

// Update messages
func Update(id, srcFile, destFile string) error {
	if len(srcFile) == 0 {
		return fmt.Errorf(i18n.Sprintf(id, "srcFile cannot be empty"))
	}

	if len(destFile) == 0 {
		return fmt.Errorf(i18n.Sprintf(id, "destFile cannot be empty"))
	}

	srcMessages, err := i18n.Unmarshal(id, srcFile)
	if err != nil {
		return err
	}
	dstMessages, err := i18n.Unmarshal(id, destFile)
	if err != nil {
		return err
	}

	result := *dstMessages
	for key, value := range *srcMessages {
		if _, ok := result[key]; !ok {
			result[key] = value
		}
	}

	var content []byte
	of := strings.ToLower(destFile)
	if strings.HasSuffix(of, ".json") {
		content, err = i18n.Marshal(result, "json")
	}
	if strings.HasSuffix(of, ".toml") {
		content, err = i18n.Marshal(result, "toml")
	}
	if strings.HasSuffix(of, ".yaml") {
		content, err = i18n.Marshal(result, "yaml")
	}
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(destFile), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destFile, content, os.ModePerm)
	if err != nil {
		return nil
	}

	fmt.Printf("Update %+v ...\n", destFile)

	return nil
}
