package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// В этой версии мы просто проходимся по разделенной строке и выделяем уникальные значения
func firstVersion(strs []string) {
	set := make(map[string]bool)

	for _, str := range strs {
		set[str] = true
	}

	fmt.Print("Unique values from string:\n[ ")
	for key := range set {
		fmt.Print(key, " ")
	}
	fmt.Println("]")
}

// В этой верии мы проходимся по разделенной строке и выделяем подмножества
func secondVersion(strs []string) {
	set := make(map[string]int)

	for _, str := range strs {
		set[str] += 1
	}

	fmt.Println("Subsets from string:")
	for key := range set {
		str := strings.Repeat(key+" ", set[key])
		fmt.Printf("[ %s] ", str)
	}
	fmt.Printf("\n")

}

func main() {
	//Вводим данные через запятую
	fmt.Println("Please enter the string set (through comma (, ))")
	inbuf := bufio.NewReader(os.Stdin)
	line, err := inbuf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//Делим строку на элементы
	strs := strings.Split(string(line[0:len(line)-1]), ", ")
	firstVersion(strs)
	secondVersion(strs)
}
