package util

import "strconv"

func StringSliceToInt(s []string) (res []int) {
	for _, x := range s {
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return
}
