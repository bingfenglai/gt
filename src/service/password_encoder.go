package service

import "golang.org/x/crypto/bcrypt"

type IPasswordEncoder interface {

	// 对密码进行加密
	Encode(src string) (string, error)

	// 检查密码是否正确
	Check(plaintext string, ciphertext string) (bool, error)
}

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
