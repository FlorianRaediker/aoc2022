package util

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func Map[T, K any](f func(T) K, items []T) (res []K) {
	for _, item := range items {
		res = append(res, f(item))
	}
	return
}

func StringToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func ToInts(s []string) (res []int) {
	for _, x := range s {
		res = append(res, StringToInt(x))
	}
	return
}

func Max[T constraints.Ordered](items ...T) (max T) {
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return
}

func Sum(items []int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return
}
