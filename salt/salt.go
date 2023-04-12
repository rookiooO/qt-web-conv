package salt

import "golang.org/x/crypto/bcrypt"

// Encrypt 将密码加密，需要传入密码返回的是加密后的密码
func Encrypt(origin string) (string, error) {
	// 加密密码，使用 bcrypt 包当中的 GenerateFromPassword 方法，bcrypt.DefaultCost 代表使用默认加密成本
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(origin), bcrypt.DefaultCost)
	if err != nil {
		// 如果有错误则返回异常，加密后的空字符串返回为空字符串，因为加密失败
		return "", err
	} else {
		// 返回加密后的密码和空异常
		return string(encryptPassword), nil
	}
}

// Equals 对比密码是否正确
func Equals(origin, salt string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	err := bcrypt.CompareHashAndPassword([]byte(salt), []byte(origin))
	// 对比密码是否正确会返回一个异常，按照官方的说法是只要异常是 nil 就证明密码正确
	return err == nil
}
