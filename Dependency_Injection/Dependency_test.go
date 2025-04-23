package dependency_test

import (
	"bytes"
	"testing"

	dependency "github.com/anicse37/TDD_Go/Dependency_Injection"
)

func TestDependency(t *testing.T) {
	buffer := &bytes.Buffer{}
	greater := dependency.Greeter{Writer: buffer}
	greater.Greet("github.com/anicse37/TDD_Go")
	want := "Hello, github.com/anicse37/TDD_Go"
	if want != buffer.String() {
		t.Errorf("Got %s || Want %s\n", buffer.String(), want)
	}
}
