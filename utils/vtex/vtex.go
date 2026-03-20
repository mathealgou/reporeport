package vtex

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reporeport/utils/types"
)

func getBuilders(manifest map[string]interface{}) []string {
	builders := []string{}

	if manifest["builders"] != nil {
		for key, val := range manifest["builders"].(map[string]interface{}) {
			str := fmt.Sprintf("%s: %v", key, val)
			builders = append(builders, str)
		}
	}

	return builders
}

func getPeerDependencies(manifest map[string]interface{}) []string {
	peerDependencies := []string{}

	if manifest["peerDependencies"] != nil {
		for key, val := range manifest["peerDependencies"].(map[string]interface{}) {
			str := fmt.Sprintf("%s: %v", key, val)
			peerDependencies = append(peerDependencies, str)
		}
	}

	return peerDependencies
}


func GetVtexCharacteristics() types.VtexCharacteristics {
	file, err := os.Open("manifest.json");

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var manifest interface{}

	err = json.Unmarshal(bytes, &manifest)

	
	result := types.VtexCharacteristics{
		Name:        manifest.(map[string]interface{})["name"].(string),
		Vendor:      manifest.(map[string]interface{})["vendor"].(string),
		Version:     manifest.(map[string]interface{})["version"].(string),
		Builders:    getBuilders(manifest.(map[string]interface{})),
		PeerDependencies: getPeerDependencies(manifest.(map[string]interface{})),
	}

	return result
}
