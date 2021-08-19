package signature

import "fmt"

// NewSignature will take some title (that can be a lib name version),
// and generate a file signature header, to be used in generated files
func NewSignature(title string) string {
	return fmt.Sprintf(signatureTmpl, title)
}
