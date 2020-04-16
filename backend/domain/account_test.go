package domain

import "testing"

func TestGenerateAccountCredentialsSuccess(t *testing.T) {
	if accountCredentials, err := GenerateAccountCredentials("info@example.com", "pass"); err != nil {
		t.Fatalf("failed test (GenerateAccountCredentials): %#v", err)
	} else if accountCredentials.Password != "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1" {
		t.Fatalf("failed test (GenerateAccountCredentials): encryptedPass is not equal")
	}

	pass := CreateLongPass(254)
	if _, err := GenerateAccountCredentials("info@example.com", pass); err != nil {
		t.Fatalf("failed test (GenerateAccountCredentials): %#v", err)
	}
}

func TestGenerateAccountCredentialsFailed(t *testing.T) {
	pass := CreateLongPass(255)
	if _, err := GenerateAccountCredentials("info@example.com", pass); err == nil {
		t.Fatalf("failed test (GenerateAccountCredentials): pass")
	}

	if _, err := GenerateAccountCredentials("info.example.com", "pass"); err == nil {
		t.Fatalf("failed test (GenerateAccountCredentials): mail")
	}

	if _, err := GenerateAccountCredentials("info@example", "pass"); err == nil {
		t.Fatalf("failed test (GenerateAccountCredentials): mail")
	}

	if _, err := GenerateAccountCredentials("@example.com", "pass"); err == nil {
		t.Fatalf("failed test (GenerateAccountCredentials): mail")
	}

	if _, err := GenerateAccountCredentials("info@", "pass"); err == nil {
		t.Fatalf("failed test (GenerateAccountCredentials): mail")
	}
}

func CreateLongPass(n int) string {
	var pass string
	for i := 0; i < n; i++ {
		pass += "a"
	}
	return pass
}
