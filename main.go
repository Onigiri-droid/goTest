package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Проверка, что передан аргумент с путем к файлу
	if len(os.Args) < 2 {
		fmt.Println("Пример: go run . War_and_Peace.txt")
		return
	}

	// Открытие файла
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Регулярное выражение для определения слов
	re := regexp.MustCompile(`[a-zA-Zа-яА-Я]+`)

	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Чтение файла построчно
	for scanner.Scan() {
		line := scanner.Text()

		// Приводим к нижнему регистру и ищем слова
		words := re.FindAllString(strings.ToLower(line), -1)

		// Увеличиваем частоту для каждого слова
		for _, word := range words {
			wordFreq[word]++
		}
	}

	// Обработка ошибок сканирования
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Создание слайса для сортировки
	type wordCount struct {
		Word  string
		Count int
	}
	var wordList []wordCount

	for word, count := range wordFreq {
		wordList = append(wordList, wordCount{Word: word, Count: count})
	}

	// Сортировка по убыванию частоты
	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].Count > wordList[j].Count
	})

	// Вывод 10 наиболее часто встречающихся слов
	for i := 0; i < 10 && i < len(wordList); i++ {
		fmt.Printf("\"%s\" mentioned %d\n", wordList[i].Word, wordList[i].Count)
	}
}
