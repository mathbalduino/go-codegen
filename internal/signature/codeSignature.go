package signature

import "fmt"

// CodeSignature will take some Library name (repo path) and
// version, and generate a file signature header, to be used in
// generated files
func CodeSignature(libNameVersion string) string {
	return fmt.Sprintf(signatureTmpl, libNameVersion)
}
