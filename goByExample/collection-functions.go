package main

import "fmt"

func Index(vs []string, t string) int{
	for i, v := range vs{
		if v == t {
			return i
		}
	}

	return -1
}

func Include(vs []string, t string) bool{
	return Index(vs, t) > 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, i := range vs {
		if !f(i){
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, i := range vs {
		if f(i){
			return true
		}
	}
	return false
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, i := range vs {
		if f(i){
			vsf = append(vsf, i)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main(){
	var strs = []string{"home", "friday", "lunch", "dinner"}
	fmt.Println(Index(strs, "friday"))

	fmt.Println(Include(strs, "lll"))

	fmt.Println(Any(strs, func(i string) bool {
		return i == "lunch"
	}))

	fmt.Println(All(strs, func(i string) bool {
		return i != "lunch"
	}))

	fmt.Println(Filter(strs, func(i string) bool {
		return i == "home"
	}))

	fmt.Println(Map(strs, func(i string) string {
		v := i + "_GM"
		return v
	}))
}

