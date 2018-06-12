package user

import "testing"

func TestHashCheckPassword(t *testing.T) {
	pass := "somerandompassword1234**&"

	hash, err := HashPassword(pass)
	if err != nil {
		t.Error(err)
	}

	match := CheckPasswordHash(pass, hash)
	if !match {
		t.Errorf("Password should be match")
	}
}
