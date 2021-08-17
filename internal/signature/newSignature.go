package signature

import "fmt"

// NewSignature will take some Library name (repo path) and
// version, and generate a file signature header, to be used in
// generated files
func NewSignature(libNameVersion string) string {
	return fmt.Sprintf(signatureTmpl, libNameVersion)
}
