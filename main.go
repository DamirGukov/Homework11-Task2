package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	MakeFile()

	file, err := os.Open("messageFile.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	regex, err := regexp.Compile(`\b[AEIOUYaeioyu][a-z\-]*`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	matches := regex.FindAllString(string(data), -1)

	var result []string
	for _, match := range matches {
		if len(match) > 1 {
			message := match
			result = append(result, message)
		}
	}

	fmt.Println("Found messages:")
	for _, message := range result {
		fmt.Println(message)
	}
}

func MakeFile() {
	file, err := os.Create("messageFile.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("It is a common thing that while living in Kyiv, she heard and saw quite often the life that her poetic instructions told about. She cannot be a bystander in the drama of national life. Such was the nature of her talent. She could not write, as if about some foreign, distant matter. She felt him near, like a beloved heart.\n\nSo, no, Lesya Ukraintsi cannot be separated from national life. Her works are incomprehensible without national content, without knowledge of national history, traditions and great artists. That is why she was destined to be an active member of her nation, to understand its needs, feel its pains, love its soul, suffer and fight for it.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
