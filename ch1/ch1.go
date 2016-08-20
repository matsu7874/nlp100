package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	p00()
	p01()
	p02()
	p03()
	p04()
	p05()
	p06()
	p07()
	p08()
	p09()
}

func p00() {
	target := "stressed"
	result := ""
	length := len(target)
	for i := length - 1; i > 0; i-- {
		result += string(target[i])
	}
	fmt.Println(result)
}

func p01() {
	target := "パタトクカシーー"
	result := ""
	for i, c := range target {
		if i%2 == 0 {
			result += string(c)
		}
	}
	fmt.Println(result)

}

func p02() {
	target := [][]rune{[]rune("パトカー"), []rune("タクシー")}
	result := ""
	length := len(target[0])
	for i := 0; i < length; i++ {
		result += string(target[0][i]) + string(target[1][i])
	}
	fmt.Println(result)
}

func p03() {
	target := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	re := regexp.MustCompile("[.,]")
	target = re.ReplaceAllString(target, "")
	var result []int
	words := strings.Split(target, " ")
	for _, w := range words {
		result = append(result, len(w))
	}
	fmt.Println(result)

}

func p04() {
	target := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	words := strings.Split(target, " ")
	uni := []int{1, 5, 6, 7, 8, 9, 15, 16}
	var result = map[string]int{}
	for i, w := range words {
		containedFlag := false
		for _, v := range uni {
			if i+1 == v {
				containedFlag = true
				break
			}
		}
		width := 2
		if containedFlag {
			width = 1
		}
		result[string(w[0:width])] = i + 1
	}
	fmt.Println(result)
}

func p05() {
	target := "I am an NLPer"
	fmt.Println(strings.Join(nGram(target, 'c', 3), ","))
	fmt.Println(strings.Join(nGram(target, 'w', 3), ","))
}
func nGram(str string, mode rune, n int) []string {
	var result []string
	if mode == 'c' {
		length := len(str)
		for i := 0; i < length-n+1; i++ {
			result = append(result, string(str[i:i+n]))
		}
	} else if mode == 'w' {
		words := strings.Split(str, " ")
		length := len(words)
		for i := 0; i < length-n+1; i++ {
			result = append(result, strings.Join(words[i:i+n], " "))
		}
	}
	return result
}

func p06() {
	aryX := nGram("paraparaparadise", 'c', 2)
	aryY := nGram("paragraph", 'c', 2)
	X := map[string]bool{}
	Y := map[string]bool{}
	for _, k := range aryX {
		X[k] = true
	}
	for _, k := range aryY {
		Y[k] = true
	}
	union := map[string]bool{}
	for k := range X {
		union[k] = true
	}
	for k := range Y {
		union[k] = true
	}
	intersection := map[string]bool{}
	for k := range X {
		if Y[k] {
			intersection[k] = true
		}
	}
	difference := map[string]bool{}
	for k := range X {
		if !Y[k] {
			difference[k] = true
		}
	}
	fmt.Println(union)
	fmt.Println(intersection)
	fmt.Println(difference)
	fmt.Println(X["se"], Y["se"])
}

func p07() {
	p07Template(12, "気温", 22.4)
}
func p07Template(x int, y string, z float64) {
	fmt.Println(strconv.Itoa(x) + "時の" + y + "は" + strconv.FormatFloat(z, 'f', 2, 64))
}

func p08() {
	target := "I couldn't believe that I could actually understand what I was reading: the phenomenal power of the human mind."
	result := ""
	for _, c := range target {
		if int('a') <= int(c) && int(c) <= int('z') {
			result += string(string(219 - int(c)))
		} else {
			result += string(c)
		}
	}
	fmt.Println(result)
	target = result
	result = ""
	for _, c := range target {
		if int('a') <= int(c) && int(c) <= int('z') {
			result += string(string(219 - int(c)))
		} else {
			result += string(c)
		}
	}
	fmt.Println(result)
}

func p09() {
	target := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	words := strings.Split(target, " ")
	var result []string
	for _, w := range words {
		length := len(w)
		if length <= 4 {
			result = append(result, string(w))
		} else {
			order := rand.Perm(length - 2)
			randomised := ""
			randomised += string(w[0])
			for _, v := range order {
				randomised += string(w[v+1])
			}
			randomised += string(w[length-1])
			result = append(result, randomised)
		}
	}
	fmt.Println(result)
}
