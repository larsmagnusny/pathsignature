package pathsignature

import (
	"testing"
)

func TestSignature_Normal(t *testing.T) {
	str := "/this/is/a/path/something.exe"

	signature := createPathSignature(str)

	signatureString := signature.toString()
	expected := "th-is-a -pa-so-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_Middle(t *testing.T) {
	str := "/this/*/a/path/something.exe"

	signature := createPathSignature(str)

	signatureString := signature.toString()
	expected := "th-**-**-**-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_Start(t *testing.T) {
	str := "*/this/is/a/path/something.exe"

	signature := createPathSignature(str)

	signatureString := signature.toString()
	expected := "**-**-**-**-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_End(t *testing.T) {
	str := "/this/is/a/path/*"

	signature := createPathSignature(str)

	signatureString := signature.toString()
	expected := "th-is-a -pa-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}
