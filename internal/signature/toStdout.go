package signature

import (
	"fmt"
	"io"
	"os"
)

// ToStdout will print the signature header to the stdout.
// If ANSI codes are supported by the stdout, the output will
// be bold green colored
func ToStdout(libNameVersion string, supportsANSI bool) {
	text := fmt.Sprintf(signatureTmpl, libNameVersion)
	if supportsANSI {
		text = terminalColorBoldGreen + text + terminalColorReset
	}

	fmt.Fprint(output, text)
}

// As a variable just to ease tests
var output io.Writer = os.Stdout

const terminalColorBoldGreen = "\033[32;1m"
const terminalColorReset = "\033[0m"
