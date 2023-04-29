package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPing(t *testing.T) {
	cmd := rootCmd
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"ping"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "Pong" {
		t.Fatalf("Expected \"%s\" got \"%s\"", "Pong", string(out))
	}
}
