package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Branch struct {
	Name   string `json:"name,omitempty"`
	Target Target `json:"target,omitempty"`
}

type Target struct {
	Hash string `json:"hash,omitempty"`
}

type BranchInput struct {
	BranchName string `json:"name"`
	BaseBranchName string `json:"baseBranchName"`
	Workspace string `json:"workspace"`
	Repository string `json:"repository"`
}

func branchRead(bi BranchInput, c *Client) (Branch, error) {
	rbReq, _ := c.Get(fmt.Sprintf("2.0/repositories/%s/%s/refs/branches/%s",
		bi.Workspace,
		bi.Repository,
		bi.BaseBranchName,
	))

	if rbReq.StatusCode == 200 {
		var b Branch

		body, readerr := ioutil.ReadAll(rbReq.Body)
		if readerr != nil {
			panic(readerr)
		}

		unmarshallErr := json.Unmarshal(body, &b)
		if unmarshallErr != nil {
			panic(unmarshallErr)
		}

		return b, nil
	}

	if rbReq.StatusCode == 404 {
		return Branch{}, fmt.Errorf("could not find branch")
	}

	return Branch{}, fmt.Errorf("could not process request, status: %d", rbReq.StatusCode)
}

func BranchCreate(bi BranchInput, c *Client) error {
	base, err := branchRead(bi, c)
	if err != nil {
		log.Println("Thee hath brought shame upon our version control")
		log.Println(err)
		return err
	}

	branch := Branch{
		Name: bi.BranchName,
		Target: Target{
			Hash: base.Target.Hash,
		},
	}

	bytedata, err := json.Marshal(branch)

	if err != nil {
		return err
	}

	_, err = c.Post(fmt.Sprintf("2.0/repositories/%s/%s/refs/branches",
		bi.Workspace,
		bi.Repository,
	), bytes.NewBuffer(bytedata))
	if err != nil {
		return err
	}

	return nil
}
