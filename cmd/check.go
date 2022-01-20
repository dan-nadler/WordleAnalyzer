/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	//"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run the checking algorithm",
	Long: `Enter each guess and the respons from Wordle. This program will tell you the number 
of possible words that remain. Note that the dictionary used by this program is not 
identical to the one used by Wordle. 

Enter 'q' to exit and print the possible remaining words.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		words := readWordleVocab()
		reader := bufio.NewReader(os.Stdin)

		for range []int{1, 2, 3, 4, 5, 6} {
			fmt.Println("Word guessed (q to quit):")
			guess := getInput(reader)
			if guess == "q" {
				break
			}

			fmt.Println("Wordle response: (g/y/x/q)")
			response := getInput(reader)
			if response == "q" {
				break
			}

			for i, r := range response {
				var filteredWords []string
				for len(words) > 0 {
					w := words[0]
					words = words[1:]

					if r == 'g' || r == 'G' {
						if w[i] != guess[i] {
							continue
						}
					}

					if r == 'y' || r == 'Y' {
						if !(strings.Contains(w, string(guess[i])) && w[i] != guess[i]) {
							continue
						}
					}

					if r == 'x' || r == 'X' {
						if strings.Contains(w, string(guess[i])) {
							continue
						}
					}

					filteredWords = append(filteredWords, w)
				}
				words = filteredWords
			}

			//rules := make([]Rule, 0)
			//rules = appendRules(word, response, rules)
			//words = eliminateWords(words, rules)
			fmt.Println(len(words), "words remaining.")
		}

		fmt.Println(words)
	},
}

func getInput(reader *bufio.Reader) string {
	word, _ := reader.ReadString('\n')
	word = strings.Replace(word, "\n", "", -1)
	word = strings.Replace(word, "\r\n", "", -1)
	return word
}

func filterVocab(vocab map[string]interface{}, numLetters int) []string {
	keys := getMapKeys(vocab)
	var filtered []string

	for i := range keys {
		if len(keys[i]) == numLetters {
			filtered = append(filtered, keys[i])
		}
	}

	return filtered
}

func getMapKeys(vocab map[string]interface{}) []string {
	keys := make([]string, 0, len(vocab))
	for key := range vocab {
		keys = append(keys, key)
	}
	return keys
}

func readWordleVocab() []string {
	file, err := os.Open("./cmd/wordle_words_2.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l[len(l)-5:])
	}
	return lines
}

func readVocab(numLetters int) []string {
	f, _ := ioutil.ReadFile("./cmd/words_dictionary.json")
	var data map[string]interface{}
	json.Unmarshal(f, &data)
	return filterVocab(data, numLetters)
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
