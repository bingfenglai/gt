package interfaces

type IPasswordEncoder interface {

	// 对密码进行加密
	Encode(src string) (string, error)

	// 检查密码是否正确
	Check(plaintext string, ciphertext string) (bool, error)
}
