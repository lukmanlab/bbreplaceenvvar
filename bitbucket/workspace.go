package bitbucket

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var WORKSPACE = os.Getenv("WORKSPACE")
var AUTH = os.Getenv("BBAUTH")
var bearer = "Basic " + base64.StdEncoding.EncodeToString([]byte(AUTH))

type Workspace struct{
	UUID	string		`json:"uuid"`
	Key		string		`json:"key"`
	Value	string		`json:"value"`
	Secured bool		`json:"secured"`
}

func GetWorkspacePipelineConfigVar() []byte {

	if (AUTH == "") || (WORKSPACE == "")  {
		fmt.Printf("Please set BBAUTH and WORKSPACE environment variables in local!\n")
		os.Exit(1)
	}

	url := "https://api.bitbucket.org/2.0/workspaces/" + WORKSPACE + "/pipelines-config/variables"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

	return body
}

func UpdateWorkspacePipelineConfigVar(key, uuid, sa string)  {

	payload := &Workspace{uuid, key, sa, false}

	url := "https://api.bitbucket.org/2.0/workspaces/"+WORKSPACE+"/pipelines-config/variables/"+uuid

    json_data, err := json.Marshal(payload)

	req, err := http.NewRequest("PUT", url, bytes.NewReader(json_data))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

	fmt.Println(string(body))
}
