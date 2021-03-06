package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
	if err := mainCore(); err != nil {
		log.Fatal(err)
	}
}

func mainCore() error {
	// Logging useful during development.
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

	// Initialize access token to be written to stdout
	execCred := &v1beta1.ExecCredential{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1beta1.SchemeGroupVersion.String(),
			Kind:       "ExecCredential",
		},
		Status: &v1beta1.ExecCredentialStatus{
			Token: "asdf1234",
		},
	}

	// If the token has an expiry then set the expirationTimestamp field as follows:
	expiry := time.Now().Add(time.Minute * 15)
	expiryK8STime := metav1.NewTime(expiry)
	execCred.Status.ExpirationTimestamp = &expiryK8STime

	// Serialize the token and write it to stdout
	execCredJSON := new(bytes.Buffer)
	if err := json.NewEncoder(execCredJSON).Encode(execCred); err != nil {
		return err
	}
	log.Debugf("execCred: %s", string(execCredJSON.Bytes()))
	if _, err := io.Copy(os.Stdout, execCredJSON); err != nil {
		return err
	}
	return nil
}
