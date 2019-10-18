//получить из файла уникальные записи по ключевому слову, первые 16 символов
package main

import (
	"bufio"
	"fmt"
	"os"
//	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {


        if len(os.Args)<2 {
		fmt.Println("Run: "+os.Args[0]+" listmpkt.csv")
        	return
	}

	fw, err := os.Create("dat.csv")
	defer fw.Close()
	check(err)


	fr, err := os.Open(os.Args[1])
	defer fr.Close()
	check(err)

	rez := ""
	row := ""

	// инициализируем набор значение ключ
	counts := make(map[string]int)

	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		row = scanner.Text()
		if len(row) > 16 {

			rez = row[0:16]
			//Увеличим на 1 значение элемента набора для ключа равного row[0:16]  
                        //( если таковой отсутствует будет дабавлен новый элемент в набор c значением равным 1
			counts[rez]++ 
		
		}
//                lineRez := strings.Split(row, ";")
//                couaddrid[lineRez[31]]++


	}
	fmt.Println("Выявлены дубликаты ШПИ:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
                        fw.WriteString(line + "\n")
			//fmt.Println(line)

	}

	// Запишем в файл все найденные adrid из реестра

//	for l, k := range couaddrid {
//      		faddr.WriteString(l+"\n") //
//                fmt.Println(k," ",l)
//
//	}


}
