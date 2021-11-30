package signature

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/internal"
	"testing"
)

func TestNewSignature(t *testing.T) {
	t.Run("Should correctly fill the signature", func(t *testing.T) {
		libNameVersion := "LibraryName v1.2.85"
		s := NewSignature(libNameVersion)
		if s != fmt.Sprintf(signatureTmpl, libNameVersion, internal.LibraryModulePath, internal.LibraryModuleVersion, internal.LibraryName) {
			t.Fatalf("Generated wrong signature")
		}
	})
}
