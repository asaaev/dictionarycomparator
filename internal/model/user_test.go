package model

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUser_HashPassword(t *testing.T) {
	user := &User{
		Email:    "test@example.com",
		Password: "12345",
	}

	err := user.HashPassword("12345")
	if err != nil {
		t.Errorf("Failed to hash password: %s", err)
	}
	if user.Password == "12345" || len(user.Password) == 0 {
		t.Errorf("Password should be hashed and not equal to the plaintext")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("12345")); err != nil {
		t.Errorf("The hash does not match the original password: %s", err)
	}

}
func TestUser_CheckPassword(t *testing.T) {
	user := &User{
		Email:    "test@examle.com",
		Password: "",
	}
	err := user.HashPassword("12345")
	if err != nil {
		t.Fatalf("Failed to hash password: %s", err)
	}
	err = user.CheckPassword("12345")
	if err != nil {
		t.Errorf("Password does not match the hash: %s", err)
	}
	err = user.CheckPassword("wrongpassword")
	if err == nil {
		t.Errorf("Password matching should fail for wrong password")
	}
}
