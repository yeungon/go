package lt_string

import (
	"fmt"
)

func Index() {
	lap := "xin chào"
	fmt.Println(lap[6]) //195
	trinh := "các bạn"
	fmt.Println(trinh[1]) //195
	fmt.Println(trinh[5]) //98

	//Chuyển sang slice rune
	laprunes := []rune(lap)
	fmt.Println(string(laprunes[6])) // "à"
	fmt.Println(string(lap[6]))      // "Ã" <- Là byte, chưa phải string

	trinhrunes := []rune(trinh)
	fmt.Println(string(trinhrunes[1])) // "á"
	fmt.Println(string(trinh[1]))      // "Ã" <- Là byte, chưa phải string

	fmt.Println(string(trinhrunes[5])) // "ạ"
	fmt.Println(string(trinh[5]))      // "b" <- Là byte, sai index, nên không đúng character "ạ"

}
