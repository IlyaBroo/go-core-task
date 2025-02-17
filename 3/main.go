package main

import "fmt"

func main() {
	testStruct := NewStringIntMap()
	testStruct.Add("key1", 1)
	testStruct.Add("key2", 2)

	fmt.Println(testStruct.Copy())

	fmt.Println("key1 is:", testStruct.Exists("key1"))
	value, exists := testStruct.Get("key1")
	fmt.Println(value, exists)

	testStruct.Remove("key1")
	value, exists = testStruct.Get("key1")
	fmt.Println(value, exists)
}

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	m := new(StringIntMap)
	m.data = make(map[string]int)
	return m
}
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
