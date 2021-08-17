package signature

import (
	"bytes"
	"fmt"
	"testing"
)

func TestToStdout(t *testing.T) {
	t.Run("If supportsANSI is false, just print the literal signature string to the output", func(t *testing.T) {
		var fakeOutput bytes.Buffer
		output = &fakeOutput
		libNameVersion := "LibraryName v5.0.8"
		ToStdout(libNameVersion, false)
		expected := fmt.Sprintf(signatureTmpl, libNameVersion)
		received := fakeOutput.String()
		if received != expected {
			t.Fatalf("Not the expected output")
		}
	})
	t.Run("If supportsANSI is true, print the signature string with the prefix/suffix ANSI bold green codes", func(t *testing.T) {
		var fakeOutput bytes.Buffer
		output = &fakeOutput
		libNameVersion := "LibraryName v5.0.8"
		ToStdout(libNameVersion, true)
		expected := terminalColorBoldGreen + fmt.Sprintf(signatureTmpl, libNameVersion) + terminalColorReset
		received := fakeOutput.String()
		if received != expected {
			t.Fatalf("Not the expected output")
		}
	})
}
