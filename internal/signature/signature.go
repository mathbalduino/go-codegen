package signature

// signatureTmpl should be used to build headers
// for generated files
const signatureTmpl = `
||
|| %s
||
|| File generated using github.com/mathbalduino/go-codegen
|| by Matheus Leonel Balduino
||
|| Everywhere, under @mathbalduino:
||   GitLab:    @mathbalduino
||   Instagram: @mathbalduino
||   Twitter:   @mathbalduino
||   WebSite:   mathbalduino.com.br/go-codegen
||
`

// FileSuffix should be used as a suffix for
// every generated file
const FileSuffix = ".mathbalduino"
