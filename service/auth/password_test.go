package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Run("should create and hash password correctly", func(t *testing.T) {
		hash, err := HashPassword("password")
		if err != nil {
			t.Errorf("error hashing password: %v", err)
		}

		if hash == "" {
			t.Error("expected hash to be not empty")
		}

		if hash == "password" {
			t.Error("expected hash to be different from password")
		}
	})

}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}

	if ComparePassword(hash, []byte("notpassword")) {
		t.Errorf("expected password to not match hash")
	}
	// t.Run("passwords should match", func(t *testing.T) {

	// })

}
