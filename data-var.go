// 基本程式結構
package main // 可執行程式必須使用 main 封包
import "fmt" // 載入內建的 fmt 封包，用來做基本的輸入輸出
func main() { // 建立 main 函式，程式的進入點
	//執行程式時，從這邊開始
	//輸出訊息到終端 fmt.Println(資料,資料)
	fmt.Println(3)      // 整數 init
	fmt.Println(3.14)   // 浮點數 float64
	fmt.Println("字串")   // 字串 string
	fmt.Println("true") // 布林
	fmt.Println('a')    // 字符 rune
}
