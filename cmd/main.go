package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type ExecCredential struct {
	APIVersion string               `json:"apiVersion"`
	Kind       string               `json:"kind"`
	Status     ExecCredentialStatus `json:"status"`
}

type ExecCredentialStatus struct {
	Token string `json:"token"`
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
	if err := mainCore(); err != nil {
		log.Fatal(err)
	}
}

func mainCore() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Debugf("wd: %s", wd)
	for i, arg := range os.Args {
		log.Debugf("args[%d]: %s", i, arg)
	}
	for i, keyValuePair := range os.Environ() {
		log.Debugf("env[%d]: %s", i, keyValuePair)
	}
	execCred := &ExecCredential{
		APIVersion: "client.authentication.k8s.io/v1beta1",
		Kind:       "ExecCredential",
		Status: ExecCredentialStatus{
			Token: "asdf1234",
		},
	}
	execCredJSONBytes, err := json.Marshal(execCred)
	if err != nil {
		return err
	}
	log.Debugf("execCred: %s", string(execCredJSONBytes))
	_, err = io.Copy(os.Stdout, bytes.NewReader(execCredJSONBytes))
	return err
}
