package signature

import (
	"fmt"
)

// ToStdout will print the signature header to the stdout.
// If ANSI codes are supported by the stdout, the output will
// be bold green colored
func ToStdout(libNameVersion string, supportsANSI bool) {
	text := fmt.Sprintf(signatureTmpl, libNameVersion)
	if supportsANSI {
		text = terminalColorGreen + text + terminalColorReset
	}

	fmt.Println(text)
}

const terminalColorGreen = "\033[32;1m"
const terminalColorReset = "\033[0m"
