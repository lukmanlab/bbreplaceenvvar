package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"bbreplaceenvvar/bitbucket"
	"os"
)

type JsonOutput struct {
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Pagelen int         `json:"pagelen"`
	Values  interface{} `json:"values"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	
	// read sa json
	saFileDev, err := os.ReadFile("./data/dev_sa.json")
    check(err)
	saFileUat, err := os.ReadFile("./data/uat_sa.json")
    check(err)
	saFileProd, err := os.ReadFile("./data/prod_sa.json")
    check(err)		

	// encode string to base64
	devSAEncoded := base64.StdEncoding.EncodeToString([]byte(string(saFileDev)))
	uatSAEncoded := base64.StdEncoding.EncodeToString([]byte(string(saFileUat)))
	prodSAEncoded := base64.StdEncoding.EncodeToString([]byte(string(saFileProd)))

	/* debug
    fmt.Println(devSAEncoded)
    fmt.Println(uatSAEncoded)
    fmt.Println(prodSAEncoded)
	os.Exit(1)
	*/

	// get data from workspace
	payload := bitbucket.GetWorkspacePipelineConfigVar()
	
	// process
	var data JsonOutput
	json.Unmarshal([]byte(payload), &data)
	for _, v := range data.Values.([]interface{}) {
		key := fmt.Sprintf("%v", (v.(map[string]interface{})["key"]))
		uuid := fmt.Sprintf("%v", (v.(map[string]interface{})["uuid"]))

		fmt.Println(key, uuid)

		if key == "GCP_SA_DEV"{
			bitbucket.UpdateWorkspacePipelineConfigVar(key,uuid,devSAEncoded)
		}else if key == "GCP_SA_UAT"{
			bitbucket.UpdateWorkspacePipelineConfigVar(key,uuid,uatSAEncoded)
		}else {
			bitbucket.UpdateWorkspacePipelineConfigVar(key,uuid,prodSAEncoded)
		}
	}
}
