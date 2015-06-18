package main

import (
	"fmt"
	"strings"
)

func letsTry() (string, string, string) {
	//string0 := "hello world"
	string1 := "good morning"
	string2 := "good afternoon"
	//string3 := "good evening"
	string4 := "AabcdefgABCDEFG"
	//var string4 string = "this also works"
	return string1, string2, string4
}

func main() {
	a, b, c := letsTry()
	//fmt.Printf("string0 is '%s',\nstring1 is '%s',\nstring2 is '%s'.\n", a, b, c)
	//const f = "%T(%v)\n"
	//fmt.Printf(f, a, a)
	//fmt.Printf(f, b, b)
	//fmt.Printf(f, c, c)
	fmt.Printf("hmmm %i, %s, %s", a[0], b[1], c[2]) //these shows the int value of the characters
	fmt.Println(c[0], c[6], c[7])
	
	fmt.Println(strings.Contains(a, "good"))
	
	if strings.Contains(a, "good") == true {
		fmt.Printf("'%v' contains the word 'good'\n", a)
	}
	
	//ContainsAny
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	//ContainsRune
	//var tester rune = 103 //should return true
	var tester rune = 10 //should return false
	another := strings.ContainsRune(a, tester)
	fmt.Println(another)
	
	//Count
	fmt.Println(strings.Count(c, "")) //returns the length of string
	fmt.Println(strings.Count(c, "a")) //returns the number of 'a'
	
	//EqualFold
	fmt.Println(strings.EqualFold("Go", "go")) //true
	fmt.Println(strings.EqualFold("GoPosA", "goposa")) //true
	fmt.Println(strings.EqualFold("우리나라", "우리나라")) //true
	fmt.Println(strings.EqualFold("우리나라", "ㅇㅜ린ㅏ라")) //false
	fmt.Println(strings.EqualFold("싸랑해사랑해", "사랑해싸랑해")) //false
	
	//Fields
	fmt.Println("Original string = '   foo bar  baz   '")
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
		//so basically this is splits function
		//lets try to make an array with products of fields
	var foo = strings.Fields("  foo   bar 	baz 	")
	fmt.Println(foo)
	
	
	
	
}
