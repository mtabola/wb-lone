package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please enter the string: ")
	inbuf := bufio.NewReader(os.Stdin)
	line, err := inbuf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//Разделяем строку
	strs := strings.Split(string(line[0:len(line)-1]), " ")
	//Применяем алгоритм для переворота строки (в данном случаем мы перемещаем элементы массива)
	for i, j := 0, len(strs)-1; i <= j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	fmt.Println("Result: ", strings.Join(strs, " "))
}
