package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/leobeosab/yeoldbrancher/bitbucket"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	inputFile := flag.String("branch-list", "", "json file containing an array of objects")
	flag.Parse()

	if _,err := os.Stat(*inputFile); os.IsNotExist(err) {
		fmt.Printf("Could not read in file: \"%s\"\n", *inputFile)
		os.Exit(1)
	}

	fileData, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var newBranches []bitbucket.BranchInput

	err = json.Unmarshal(fileData, &newBranches)
	if err != nil {
		fmt.Println(err)
	}

	client := &bitbucket.Client{
		Username: os.Getenv("BITBUCKET_USERNAME"),
		Password: os.Getenv("BITBUCKET_PASSWORD"),
		HTTPClient: &http.Client{},
	}

	for _, bi := range newBranches {
		berr := bitbucket.BranchCreate(bi, client)
		if berr != nil {
			fmt.Println(berr)
		}
	}
}