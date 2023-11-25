package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please enter your expression (like: a + b): ")
	inbuf := bufio.NewReader(os.Stdin)
	line, err := inbuf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	members := strings.Split(line[:len(line)-1], " ")
	if len(members) != 3 {
		fmt.Println("Wrong input!!! Please enter like: a + b")
		os.Exit(1)
	}
	// Создаем 3 переменные отвечающие за 1 и 2 число, а также результат
	// Использумем пакет big для того, чтобы хранить огромные занчения (примерно 17 квантилионнов)
	firNum := new(big.Float)
	secNum := new(big.Float)
	result := new(big.Float)

	// Заносим данные в переменные
	firNum.SetString(members[0])
	secNum.SetString(members[2])

	// Реализумем математические операции по средством методов, которые заложены в объектах
	switch members[1] {
	case "+":
		result.Add(firNum, secNum)
	case "-":
		result.Sub(firNum, secNum)
	case "*":
		result.Mul(firNum, secNum)
	case "/":
		result.Quo(firNum, secNum)
	default:
		fmt.Println(`Unexpected operator!!! You can use only:
- Add [+]
- Sub [-]
- Mul [*]
- Div [/]`)
		os.Exit(1)
	}
	fmt.Println("Result = ", result.Text('f', 10))
}
