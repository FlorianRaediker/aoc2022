package util

import (
	"aoc2022/secrets"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetInput(year, day int) string {
	filepath := fmt.Sprintf("%v/day%v/input", year, day)

	var input string

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			panic(err)
		}

		cookie := new(http.Cookie)
		cookie.Name, cookie.Value = "session", secrets.Session
		req.AddCookie(cookie)

		client := http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			panic(fmt.Sprintf("could not get input for year %v day %v: %v: %v", year, day, resp.Status, string(body)))
		}

		input = string(body)

		err = os.WriteFile(filepath, []byte(input), 0755)

		if err != nil {
			panic(err)
		}
	} else {
		file, err := os.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		input = string(file)
	}

	return strings.TrimRight(input, "\n")
}

func GetInputStrings(year, day int) []string {
	return strings.Split(GetInput(year, day), "\n")
}

func GetInputInts(year, day int) (res []int) {
	return StringSliceToInt(GetInputStrings(year, day))
}

func GetInputBinary(year, day int) (bitlen int, res []int) {
	for _, x := range GetInputStrings(year, day) {
		i, err := strconv.ParseInt(x, 2, 0)
		if bitlen == 0 {
			bitlen = len(x)
		} else if bitlen != len(x) {
			panic("not all numbers have the same bit length")
		}
		if err != nil {
			panic(err)
		}
		res = append(res, int(i))
	}
	return
}
