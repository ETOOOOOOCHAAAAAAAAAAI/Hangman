package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func loadWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := strings.TrimSpace(scanner.Text())
		w = strings.ToLower(w)
		words = append(words, w)
	}
	return words, nil
}

func takeRandomWord(words []string) ([]rune, string) {
	idx := rand.Intn(len(words) - 1)
	secret := words[idx]
	secretRunes := []rune(secret)
	return secretRunes, secret

}
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
	words, err := loadWords("words.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла.", err)
	}
	var wrong int
	var MaxWrong = 7
	rand.Seed(time.Now().UnixNano())
	randomWordInRunes, secretWord := takeRandomWord(words)
	shown := make([]rune, len(randomWordInRunes))
	for i := range shown {
		shown[i] = '_'
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Игра начилась")
	fmt.Println(string(shown))
	for {
		fmt.Println("У вас осталось", MaxWrong-wrong, "возможности ошибиться.")
		scanner.Scan()
		s := scanner.Text()
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		rs := []rune(s)
		ch := rs[0]
		if len(rs) != 1 || !unicode.IsLetter(ch) {
			fmt.Println("Пишите нормально.")
			continue
		}
		found := false
		for i := range randomWordInRunes {
			if randomWordInRunes[i] == ch {
				shown[i] = ch
				found = true
			}
		}
		if !found {
			fmt.Println(stages[wrong])
			wrong++
		} else {
			fmt.Println("Есть такая буква!!")
			fmt.Println(string(shown))
		}
		if string(shown) == string(secretWord) {
			fmt.Println("Поздравляю вы выиграли!Слово было:")
			fmt.Println(string(secretWord))
			break
		} else if wrong == 7 {
			fmt.Println("Ты проиграл.Слово было:")
			fmt.Println(string(secretWord))
			break
		}

	}
}
