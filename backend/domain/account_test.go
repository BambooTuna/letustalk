package domain

import "testing"

func TestNewAccountCredentialsSuccess(t *testing.T) {
	if accountCredentials, err := NewAccountCredentials("info@example.com", "pass"); err != nil {
		t.Fatalf("failed test (NewAccountCredentials): %#v", err)
	} else if accountCredentials.Password != "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1" {
		t.Fatalf("failed test (NewAccountCredentials): encryptedPass is not equal")
	} else if accountCredentials.Accessible("pass") != true {
		t.Fatalf("failed test (NewAccountCredentials): Inaccessible")
	} else if accountCredentials.Position != General {
		t.Fatalf("failed test (NewAccountCredentials): Default Position should be General")
	} else if accountCredentials.Activated != false {
		t.Fatalf("failed test (NewAccountCredentials): Default Activated should be false")
	} else if accountCredentials.ChangePosition(Mentor).Position != Mentor {
		t.Fatalf("failed test (NewAccountCredentials): ChangePosition failed")
	}

	pass := CreateLongPass(254)
	if _, err := NewAccountCredentials("info@example.com", pass); err != nil {
		t.Fatalf("failed test (NewAccountCredentials): %#v", err)
	}
}

func TestNewAccountCredentialsFailed(t *testing.T) {
	pass := CreateLongPass(255)
	if _, err := NewAccountCredentials("info@example.com", pass); err == nil {
		t.Fatalf("failed test (NewAccountCredentials): pass")
	}

	if _, err := NewAccountCredentials("info.example.com", "pass"); err == nil {
		t.Fatalf("failed test (NewAccountCredentials): mail")
	}

	if _, err := NewAccountCredentials("info@example", "pass"); err == nil {
		t.Fatalf("failed test (NewAccountCredentials): mail")
	}

	if _, err := NewAccountCredentials("@example.com", "pass"); err == nil {
		t.Fatalf("failed test (NewAccountCredentials): mail")
	}

	if _, err := NewAccountCredentials("info@", "pass"); err == nil {
		t.Fatalf("failed test (NewAccountCredentials): mail")
	}
}

func CreateLongPass(n int) string {
	var pass string
	for i := 0; i < n; i++ {
		pass += "a"
	}
	return pass
}
