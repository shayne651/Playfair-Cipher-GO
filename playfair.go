package main

import ("fmt"
		"strings")

func main() {
	keyword :="keyword"
	letters := initializeDataSet(keyword)
	fmt.Println(letters[0][0])
}

func initializeDataSet(keyword string) [5][5]string {
	x :=0
	pos := 1
	alpos :=1
	c :=0
	alph := "abcdefghijklmnoprstuvwxyz"
	letters := [5][5]string{}

	for c<=len(keyword)-1{
		if strings.Contains(alph , (string([]rune(keyword)[c]))){
			alph = strings.Replace(alph , (string([]rune(keyword)[c])) , "" , -1)
		}
		c++
	}

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
	return letters
}