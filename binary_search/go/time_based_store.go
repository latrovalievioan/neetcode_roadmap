package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Println(obj.Get("love", 5))
	fmt.Println(obj.Get("love", 10))
	fmt.Println(obj.Get("love", 15))
	fmt.Println(obj.Get("love", 20))
	fmt.Println(obj.Get("love", 25))
}

type P struct {
	v string
	t int
}

type TimeMap struct {
	KeyMap map[string][]P
}

func Constructor() TimeMap {
	return TimeMap{
		KeyMap: make(map[string][]P),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	v, ok := m.KeyMap[key]
	if ok {
		m.KeyMap[key] = append(v, P{v: value, t: timestamp})
		return
	}

	m.KeyMap[key] = []P{{v: value, t: timestamp}}
}

func (m *TimeMap) Get(key string, timestamp int) string {
	v, ok := m.KeyMap[key]

	if !ok {
		return ""
	}

	return search(v, timestamp)
}

func search(xs []P, timestamp int) string {
	leftP, rightP := 0, len(xs)-1

	for {
		mid := (leftP + rightP) / 2
		current := xs[mid]

		if leftP > rightP {
			if current.t <= timestamp {
				return current.v
			} else {
				return ""
			}
		}

		if current.t > timestamp {
			rightP = mid - 1
		} else {
			leftP = mid + 1
		}
	}
}
