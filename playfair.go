package main

import "fmt"

func main() {
	keyword :="keyword"
	x :=0
	pos := 1
	alpos :=1
	alph := "abcdefghijklmnoprstuvwxyz"
	letters := [5][5]string{}

	for x<=4{
		z:=0
		for z<=4{
			if len(keyword)+1 <= pos {
					letters[x][z] = (string([]rune(alph)[alpos-1]))
					fmt.Println(letters[x][z])
				alpos++
			}else{
					letters[x][z] = (string([]rune(keyword)[pos-1]))
					fmt.Println(letters[x][z])
				pos++
			}
			z++
		}
		x++
	}
}