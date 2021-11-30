package signature

// signatureTmpl should be used to build headers
// for generated files
const signatureTmpl = `
||
|| %s
||
|| File generated using %s %s
|| by Matheus Leonel Balduino
||
|| Everywhere, under @mathbalduino:
||   GitLab:    @mathbalduino
||   Instagram: @mathbalduino
||   Twitter:   @mathbalduino
||   WebSite:   mathbalduino.com.br/%s
||
`

// FileSuffix should be used as a suffix for
// every generated file
const FileSuffix = ".mathbalduino"
