package signature

import "fmt"

// CodeSignature will take some Library name (repo path) and
// version, and generate a file signature header, to be used in
// generated files
func CodeSignature(libNameVersion string) string {
	return fmt.Sprintf(signatureTmpl, libNameVersion)
}

const signatureTmpl = `
||
|| %s
|| by Matheus Leonel Balduino
||
|| GitLab:    @matheuss_balduino
|| Instagram: @matheuss_balduino
|| WebSite:   matheus-balduino.com.br
||
`
