package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var stages = []string{
		`
	  +---+
	  |   |
	      |
	      |
	      |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	      |
	      |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	  |   |
	      |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	 /|   |
	      |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	 /|\  |
	      |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	 /|\  |
	 /    |
	      |
	=========`,
		`
	  +---+
	  |   |
	  O   |
	 /|\  |
	 / \  |
	      |
	=========`,
	}
	words := []string{"грузчик", "автомат", "машина", "зиберт", "анатомия", "женщина", "мужчина", "кулак"}
	var MaxWrong int
	var wrong = 7
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(words))
	secret := words[idx]
	secretRunes := []rune(secret)
	shown := make([]rune, len(secretRunes))
	for i := range shown {
		shown[i] = '_'
	}
	fmt.Println("Игра начилась")
	for {
		fmt.Println("У вас осталось", wrong-MaxWrong, "возможности ошибиться.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		rs := []rune(s)
		if len(rs) != 1 {
			fmt.Println("Пишите нормально.")
			continue
		}
		ch := rs[0]
		found := false
		for i := range secretRunes {
			if secretRunes[i] == ch {
				shown[i] = ch
				found = true
			}
		}
		if !found {
			fmt.Println(stages[MaxWrong])
			MaxWrong++
		} else {
			fmt.Println("Есть такая буква!!")
			fmt.Println(string(shown))
		}
		if string(shown) == string(secret) {
			fmt.Println("Поздравляю вы выиграли!Слово было:")
			fmt.Println(string(secret))
			break
		} else if MaxWrong == 7 {
			fmt.Println("Ты проиграл.Слово было:")
			fmt.Println(string(secret))
			break
		}

	}
}
