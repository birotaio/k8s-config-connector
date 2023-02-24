package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	VERSION_FILE_PATH           = "../operator/channels/stable"
	CRDS_FILE_PATH_TEMPLATE     = "../operator/channels/packages/configconnector/%s/crds.yaml"
	CRDS_OLD_FILE_PATH_TEMPLATE = "../operator/channels/packages/configconnector/%s/crds.old.yaml"
)

type VersionManifest struct {
	Manifests [](map[string]string)
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("You should provide an argument with comma separated values")
	}

	versionFilePath := path.Join(currentDir, VERSION_FILE_PATH)
	versionFileContent, err := os.ReadFile(versionFilePath)
	if err != nil {
		log.Fatal(err)
	}

	versionManifest := &VersionManifest{}
	err = yaml.Unmarshal(versionFileContent, versionManifest)
	if err != nil {
		log.Fatal(err)
	}

	version := versionManifest.Manifests[0]["version"]

	println("Version", version)

	err = editCrdFile(strings.Split(os.Args[1], ","), version, currentDir)
	if err != nil {
		log.Fatal(err)
	}
}

func editCrdFile(whitelist []string, version string, currentDir string) error {
	crdsFilePath := path.Join(currentDir, fmt.Sprintf(CRDS_FILE_PATH_TEMPLATE, version))
	crdsContent, err := os.ReadFile(crdsFilePath)
	if err != nil {
		return err
	}
	if strings.HasPrefix(string(crdsContent), "#fixed-at:") {
		timeUnixString, _ := strings.CutPrefix(strings.Split(string(crdsContent), "\n")[0], "#fixed-at:")
		timeUnix, err := strconv.Atoi(timeUnixString)
		if err != nil {
			return err
		}
		return fmt.Errorf("this manifest was already fixed at %s.\nRestore previous crds.old or from remote before fixing", time.Unix(int64(timeUnix), 0).Local().String())
	}

	out := []string{}
	found := []string{}

	for _, crd := range strings.Split(string(crdsContent), "---\n") {
		crdContentDecoded := map[interface{}]interface{}{}
		err = yaml.Unmarshal([]byte(crd), crdContentDecoded)
		if err != nil {
			return err
		}
		kind := (crdContentDecoded["spec"]).(map[interface{}]interface{})["names"].(map[interface{}]interface{})["kind"].(string)
		if contains(whitelist, kind) {
			out = append(out, crd)
			found = append(found, kind)
		}
	}

	if len(whitelist) != len(out) {
		return fmt.Errorf("did not find matches for all crds provided.\n  Requested: %s\n  Got: %s", strings.Join(whitelist, ", "), strings.Join(found, ", "))
	}

	crdsOldFilePath := path.Join(currentDir, fmt.Sprintf(CRDS_OLD_FILE_PATH_TEMPLATE, version))
	if _, err := os.Stat(crdsOldFilePath); !os.IsNotExist(err) {
		err = os.Remove(crdsOldFilePath)
		if err != nil {
			return err
		}
	}
	err = os.WriteFile(crdsOldFilePath, crdsContent, 0777)
	if err != nil {
		return err
	}

	err = os.WriteFile(crdsFilePath, []byte(fmt.Sprintf("#fixed-at:%d\n", time.Now().Unix())+strings.Join(out, "---\n")), 0)
	if err != nil {
		return err
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
