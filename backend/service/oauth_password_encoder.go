package service

import "golang.org/x/crypto/bcrypt"

type passwordEncoder struct {
}

func (p *passwordEncoder) Encode(src string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(password), nil
}

func (p *passwordEncoder) Check(plaintext string, ciphertext string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))

	return err == nil, err
}
