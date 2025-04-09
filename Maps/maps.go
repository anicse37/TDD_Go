package maps

import (
	"errors"
	"reflect"
)

type MyMaps map[int]string

var (
<<<<<<< HEAD
	ErrorAlreadyExist = errors.New("This key,value pair already exists")
	ErrorNotFound     = errors.New("The number was not found")
=======
	ErrorAlreadyExist = errors.New("value already exists for the key")
	ErrorNotFound     = errors.New("key not found in map")
>>>>>>> branch-1
)

func Mapfunction(text string, map1 map[int]string) bool {
	for _, value := range map1 {
		if text == value {
			return true
		}
	}
	return false
}

func (m MyMaps) AddValue(key int, value string) (MyMaps, error) {
<<<<<<< HEAD
	for index, _ := range m {
		if m[index] == value && index == key {
=======
	for k, v := range m {
		if k == key && v == value {
>>>>>>> branch-1
			return m, ErrorAlreadyExist
		}
	}
	m[key] = value
	return m, nil
}
<<<<<<< HEAD
func AreMapsEqual(a, b map[int]string) bool {
	return reflect.DeepEqual(a, b)
}
func (m MyMaps) Found(str string) bool {
	for index, _ := range m {
		if m[index] == str {
			return true
		}
	}
	return false
}

func (m MyMaps) Delete(key int) (MyMaps, error) {
	newMap := MyMaps{}
	for k, _ := range m {
		if k != key {
			newMap[k] = m[k]
		}
	}
	if AreMapsEqual(m, newMap) {
		return m, ErrorNotFound
	}
	return newMap, nil
=======

func (m MyMaps) Delete(key int) (MyMaps, error) {
	if _, exists := m[key]; !exists {
		return m, ErrorNotFound
	}
	delete(m, key)
	return m, nil
}

func (m MyMaps) Found(value string) (bool, error) {
	for _, v := range m {
		if v == value {
			return true, nil
		}
	}
	return false, ErrorNotFound
}

// Compares two MyMaps for equality
func AreMapsEqual(a, b MyMaps) bool {
	return reflect.DeepEqual(a, b)
>>>>>>> branch-1
}
