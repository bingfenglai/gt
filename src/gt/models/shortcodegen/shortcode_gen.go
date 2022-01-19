package shortcodegen

// 链接长短转换接口
type ShortCodeGenerator interface {

	// 将长链接转换为4组code
	GenShortCode(link string) ([]string, error)

	// 获取当前生成短码的方式
	GetGenMethod() string
}

// const chars := []string{"q","w","e","r","t","y","u","i","o","p",
// "a","s","d","f","g","h","j","k","l","z",
// "x","c","v","b","n","m","1","2","3","4",
// "5","6","7","8","9","0","Q","W","E","R",
// "T","Y","U","I","O","P","A","S","D","F",
// "G","H","J","K","L","Z","X","C","V","B",
// "N","M"}

// code 待选字符
func chars() []string {

	chars := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
		"a", "s", "d", "f", "g", "h", "j", "k", "l", "z",
		"x", "c", "v", "b", "n", "m", "1", "2", "3", "4",
		"5", "6", "7", "8", "9", "0", "Q", "W", "E", "R",
		"T", "Y", "U", "I", "O", "P", "A", "S", "D", "F",
		"G", "H", "J", "K", "L", "Z", "X", "C", "V", "B",
		"N", "M"}

	return chars
}
