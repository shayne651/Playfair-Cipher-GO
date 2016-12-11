package main

import ("fmt"
		"strings"
		"math")

func main() {
	keyword :="playfir"
	letters := initializeDataSet(keyword)
	fmt.Println(splitLetters(letters,"osmosis",true))
	fmt.Println(splitLetters(letters,"wxuwmctz",false))
}

func initializeDataSet(keyword string) [5][5]string {
	x :=0
	pos := 1
	alpos :=1
	c :=0
	alph := "abcdefghijklmnoprstuvwxyz"
	letters := [5][5]string{}

	q:=1
	for q<=len(keyword)-1{
			substring := keyword[q:len(keyword)]
			holder:= keyword[0:q+1]
			substring = strings.Replace(substring ,(string([]rune(keyword)[q])) ,"" , -1)
			keyword = holder + substring
			q++
	}

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
				alpos++
			}else{
					letters[x][z] = (string([]rune(keyword)[pos-1]))
				pos++
			}
			z++
		}
		x++
	}
	return letters
}

func encrypt(data [5][5]string, letters string) string{
	l1x:=0
	l2x:=0
	l1y:=0
	l2y:=0
	found := false
	l1 := (string([]rune(letters)[0]))
	l2 := (string([]rune(letters)[1]))
	for l1x <=4 {
		l1y=0
		for l1y<=4 {
			if l1 == data[l1x][l1y]{
				found = true
				break
			}
			l1y++
		}
		if found == true {
			found = false
			break
		}
		l1x++
	}

	for l2x <=4{
		l2y =0
		for l2y <=4{
			if l2 == data[l2x][l2y]{
				found = true 
				break
			}
			l2y++
		}
		if (found == true){
			break 
		}
		l2x++
	}

	//case 1
	if l2x == l1x {
		l2x++
		l1x++
		if l2x > 4 {
			l2x =0
		}
		if l1x > 4{
			l1x =0
		}
	}
	//case 2
	if l2y == l1y {
		l2y++
		l1y++
		if l1y > 4{
			l1y =0
		}
		if l2y > 4{
			l2y =0
		}
	}
	//case 3
	if l2y != l1y && l2x != l1x{
		holder := l1y
		l1y = l2y
		l2y = holder
	}
	returnVal := data[l1x][l1y]+data[l2x][l2y]
	return (returnVal)
}

func unEncrypt(data [5][5]string, letters string) string{
	l1x:=0
	l2x:=0
	l1y:=0
	l2y:=0
	found := false
	l1 := (string([]rune(letters)[0]))
	l2 := (string([]rune(letters)[1]))
	for l1x <=4 {
		l1y=0
		for l1y<=4 {
			if l1 == data[l1x][l1y]{
				found = true
				break
			}
			l1y++
		}
		if found == true {
			found = false
			break
		}
		l1x++
	}

	for l2x <=4{
		l2y =0
		for l2y <=4{
			if l2 == data[l2x][l2y]{
				found = true 
				break
			}
			l2y++
		}
		if (found == true){
			break 
		}
		l2x++
	}

	//case 1
	if l2x == l1x {
		if l2x >0 {
			l2x--
		}else{
			l2x=4
		}
		if l1x > 0 {
			l1x--
		}else{
			l1x=4
		}
	}
	//case 2
	if l2y == l1y {
		if l2y >0 {
			l2y--
		}else{
			l2y=4
		}
		if l1y > 0 {
			l1y--
		}else{
			l1y=4
		}
	}
	//case 3
	if l2y != l1y && l2x != l1x{
		holder := l1y
		l1y = l2y
		l2y = holder
	}
	returnVal:=data[l1x][l1y]+data[l2x][l2y]
	return (returnVal)
}
func splitLetters(data [5][5]string , keyword string , encryption bool) string{
	returnVal := ""
	z:=0 
	leng:=float64(len(keyword))
	if math.Mod(leng,2)!=0{
		keyword+="x"
		fmt.Println(keyword)
	}
	for z <=len(keyword)-1{
		if (encryption ==true){
			returnVal+=encrypt(data , keyword[z:z+2])
			fmt.Println(keyword[z:z+2], " ", returnVal)
			z+=2
			}else{
			returnVal+=unEncrypt(data , keyword[z:z+2])
			fmt.Println(keyword[z:z+2], " ", returnVal)
			z+=2
			}
	}
	return returnVal
}