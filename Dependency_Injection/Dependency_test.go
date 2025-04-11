package dependency_test

import (
	dependency "aniket/Dependency_Injection"
	"bytes"
	"testing"
)

func TestDependency(t *testing.T) {
	buffer := &bytes.Buffer{}
	greater := dependency.Greeter{Writer: buffer}
	greater.Greet("Aniket")
	want := "Hello, Aniket"
	if want != buffer.String() {
		t.Errorf("Got %s || Want %s\n", buffer.String(), want)
	}
}
