package pathsignature

import (
	"testing"
)

func TestSignature_Normal(t *testing.T) {
	str := "/this/is/a/path/something.exe"

	signature := Create(str)

	signatureString := signature.ToString()
	expected := "th-is-a -pa-so-ex-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Reverse_Normal(t *testing.T) {
	str := "/this/is/a/path/something.exe"

	signature := CreateReverse(str)

	signatureString := signature.ToString()
	expected := "ex-so-pa-a -is-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_Middle(t *testing.T) {
	str := "/this/*/a/path/something.exe"

	signature := Create(str)

	signatureString := signature.ToString()
	expected := "th-**-**-**-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Reverse_Wildcard_Middle(t *testing.T) {
	str := "/this/*/a/path/something.exe"

	signature := CreateReverse(str)

	signatureString := signature.ToString()
	expected := "ex-so-pa-a -**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_Start(t *testing.T) {
	str := "*/this/is/a/path/something.exe"

	signature := Create(str)

	signatureString := signature.ToString()
	expected := "**-**-**-**-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Reverse_Wildcard_Start(t *testing.T) {
	str := "*/this/is/a/path/something.exe"

	signature := CreateReverse(str)

	signatureString := signature.ToString()
	expected := "ex-so-pa-a -is-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Wildcard_End(t *testing.T) {
	str := "/this/is/a/path/*"

	signature := Create(str)

	signatureString := signature.ToString()
	expected := "th-is-a -pa-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}

func TestSignature_Reverse_Wildcard_End(t *testing.T) {
	str := "/this/is/a/path/*"

	signature := CreateReverse(str)

	signatureString := signature.ToString()
	expected := "**-**-**-**-**-**-**-**"
	if signatureString != expected {
		t.Errorf("Signature %s not equal to expected %s", signatureString, expected)
	}
}
