# Ответы на вопросы
## 1. Какой самый эффективный способ конкатенации строк?
**Ответ:** существуют разные способы конкатенации строк, но из них, по эффективности, выделяются следующие:
- Структура `bytes.Buffer` - так как онa работает с массивом байтов, а не со строками, что значительно выигрывает в производительности из-за того, что нет лишних конвертаций.
Пример:
```golang
package main
 
import (
    "fmt"
    "bytes"
)
 
func main() {
    var b bytes.Buffer
     
    b.WriteString("abc")
    b.WriteString("def") // добавляем
     
    fmt.Println(b.String()) // abcdef
}
```
- Структура `strings.Builder` - использует наименьшее колличество памяти для объединения строк. Также WriteString использует более быстрый способ объединения строк.
```golang
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    var sb strings.Builder
    sb.WriteString("First")
    sb.WriteString("Second")
    fmt.Println(sb.String())    // FirstSecond
}
```
## 2. Что такое интерфейсы, как они применяются в Go?
**Ответ:** интерфейс - это абстракция поведения других типов. Он позволяет определять методы, которые не привязаны к какой либо реализации. С помощью них можно реалзиовывать разные группы методов и применять их к разным объектам. В Golang они применяются для того, чтобы реализовывать **полиморфизм** (как показано в примере 21)
Также преимущество использования интерфейсов заключается в том, что они позволяют писать гибкий код, который может работать с различными типами данных, при условии, что они удовлетворяют требованиям интерфейса.

## 3. Чем отличаются RWMutex от Mutex?
**Ответ:** главное отличие `Mutex` от `RWMutex` это в том он становится более медленным, если дело доходит до чтения-записи из одной области памяти, т.к `Mutex` блокирует область памяти полностью. `RWMutex` же обладает отдельными методами, которые блокируют либо запись, либо чтение.
```golang
//Mutex
var mu sync.Mutex
...
mu.Lock()// Блокировка и разблокировка области данных
mu.Unlock()
```
vs
```golang
//RWMutex
var mu sync.RWMutex
...
mu.RLock() // Блокировка и разблокировка на запись
defer mu.Unlock()
...
mu.RLock() // Блокировка и разблокировка на чтение
defer mu.RUnlock()
```
## 4. Чем отличаются буферизированные и не буферизированные каналы?
**Ответ:** тем, что в небуферизированный канал могут принимать неограниченное количество данных из горутин. Буферизированные каналы, в свою же очередь, имеют определенный размер и блокируются, если значения в нем достигли своего потолка. Также их различие заключается в том, что в буфферезированный канал обеспечивает синхронную передачу данных между горутинами и гарантирует взаимную блокировку до завершения передачи, а не буфферезированный, наоборот, предназначен для асинхронного отправления и приниема значений, если есть свободное место в буфере.
Достигают они этого с помощью своих особенностей:
|Буферизированный |Небуферизированный |
|-----------------|-------------------|
|Отправка значения в буферизированный канал не блокируется, если буфер еще не заполнен|Отправка значения в небуферизированный канал блокируется до тех пор, пока значение не будет прочитано из канала|
|Прием значения из буферизированного канала также не блокируется, если в буфере есть значения для чтения|Прием значения из небуферизированного канала также блокируется, если в канале нет отправленных значений|


## 5. Какой размер у структуры struct{}{}?
**Ответ: 0 байт**, так как это составной литирал, который создает значение типа struct{}. Так как наша структура и так пуста, соответственно и сам литерал тоже будет пуст.
## 6. Есть ли в Go перегрузка методов или операторов?
**Ответ:** В Golang, к сожалению, не предусмотрены перегрузки методов и операторов.
## 7. В какой последовательности будут выведены элементы map[int]int?
**Ответ:** все зависит от того как мы будем выводить. Если будем выводить через `fmt.Printf("%v\n", m)`, то получим элементы в порядке возрастания индексов. Если же мы будем проходиться по мапе циклом, то будет выводиться в порядке очередности захода.
```golang
func main() {
	m := map[int]int{}
	m[1] = 128
	m[0] = 1
	m[2] = 256

	for _, v := range m {
		fmt.Print(v)
	}
}
```
`out: 128 1 256`
## 8. В чем разница make и new?
**Ответ:** 
- `make` нужен для того, чтобы выделять память из кучи для *slice*, *map* и *channel*. Также `make` не просто выделяет память, но и инициализирует ее, возвращая при этом не указатель а само значение
- `new` возвращает указатель на неинициализированное значение из кучи. Подходит для всех типов данных.
## 9. Сколько существует способов задать переменную типа slice или map?
**Ответ:** я заню порядка 4 способов, как можно создать map, и порядка 3, как slice:
<table>
<tr>
<td>Map</td>
</tr>
<tr>
<td>

```golang
        var m map[string]int // nil map

        m1 := make(map[string]float64) // Пустая, но инициализированная map
        m2 := make(map[string]float64, 100) // Подготовленная map на 100 эл-тов

        m3 := map[string]float64{// Map литерал
            "e":  2.71828,
            "pi": 3.1416,
        }
```
</td>
</tr>
<tr>
<td>Slice</td>
</tr>
<tr>
<td>

```golang
        var s []int // nil slice

        s2 := make([]float64, 100) // Подготовленный slice на 100 эл-тов

        s3 := []int8{1, 2, 3, 4, 5} // slice литерал
```

</td>
</tr>
</table>

## 10. Что выведет данная программа и почему?
```golang
func update(p *int) {
    b := 2
    p = &b
}

func main() {
    var (
        a = 1
        p = &a
    )
    fmt.Println(*p)
    update(p)
    fmt.Println(*p)
}
```
**Ответ:** `1 1`, т.к при вызове функции `update(p)`, внутри функции изменяется значение указателя p на адрес переменной b но, это изменение происходит только внутри функции, и не влияет на значение указателя p в функции main.

## 11. Что выведет данная программа и почему?
```golang
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg sync.WaitGroup, i int) {
            fmt.Println(i)
            wg.Done()
        }(wg, i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```
**Ответ:** все значения `fmt.Println(i)` и **deadlock!**, так как в горутину мы передаем не ссылку на WaitGroup, а ее копию.
## 12. Что выведет данная программа и почему? 
```golang
func main() {
    n := 0
    if true {
        n := 1
        n++
    }
    fmt.Println(n)
}
```
**Ответ:** `0` т.к это две разные `n`, у которых разные области видимости. При инкрименте мы обращаемся к `n`,которая находится в `if true{}` области видимости.
## 13. Что выведет данная программа и почему? 
```golang
func someAction(v []int8, b int8) {
    v[0] = 100
    v = append(v, b)
}

func main() {
    var a = []int8{1, 2, 3, 4, 5}
    someAction(a, 6)
    fmt.Println(a)
}
```
**Ответ:** `[100, 2, 3, 4, 5]`, т.к мы работаем с копией слайса, а не с указателем. Почему же тогда 1 элемент 100?
Потому что, слайс - это есть ссылка на массив, который был создан под копотом. Когда мы обращаемся таким образом: `v[0] = 100`, то мы говорим компилятору следующее: обратись по адресу этого массива, в ячейку 0, где нужно изменить данные.
## 14. Что выведет данная программа и почему? 
```golang
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
**Ответ:** `[b b a][a a]` практически такая же ситуация, как и с прошлым вопросом, только сначала мы вызвали функцию `append()`. И из-за того, что мы ее вызвали, у нас появилась копия массива, которая находится в другой области памяти, поэтому можно считать, что мы работали не с оригинальным слайсом, а с его копией.