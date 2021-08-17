package signature

import (
	"fmt"
	"testing"
)

func TestNewSignature(t *testing.T) {
	t.Run("Should correctly fill the signature", func(t *testing.T) {
		libNameVersion := "LibraryName v1.2.85"
		s := NewSignature(libNameVersion)
		if s != fmt.Sprintf(signatureTmpl, libNameVersion) {
			t.Fatalf("Generated wrong signature")
		}
	})
}
