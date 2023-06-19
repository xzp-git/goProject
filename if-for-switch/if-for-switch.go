package main

import "fmt"

func main() {

	score := 70

	if score > 80 {
		fmt.Println("良好")
	} else if score > 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("一般")
	}

	str := "tiantian向上"

	runeStr := []rune(str)

	for i := 0; i < len(runeStr); i++ {
		fmt.Printf("%c\r\n", runeStr[i])
	}

	for key, value := range runeStr {
		fmt.Printf("%d %c\r\n", key, value)
	}

	switch {
	case score > 80:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if i == 0 && j == 3 {
				goto over
			}

		}
	}

over:
	fmt.Println("over")
}
