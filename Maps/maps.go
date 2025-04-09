package maps

import (
	"errors"
)

var (
	Error_NotFound          = errors.New("Value Not Found")
	Error_FileAlreadyExists = errors.New("This Key Value Pair Already Exists")
)

type MyMaps map[int]string

func (m MyMaps) TestMap() MyMaps {
	return m

}

/*---------------------------------------------------*/
func AreMapsSame(got, want MyMaps) bool {
	if len(got) != len(want) {
		return false
	}
	for i, v := range want {
		if got[i] != v {
			return false
		}
	}
	return true
}

/*---------------------------------------------------*/
func (m MyMaps) Search(value string) (bool, error) {
	for i, _ := range m {
		if m[i] == value {
			return true, nil
		}
	}
	return false, Error_NotFound
}

/*---------------------------------------------------*/
func (m MyMaps) Add(key int, value string) (MyMaps, error) {
	for i, _ := range m {
		if i == key && m[i] == value {
			return m, Error_FileAlreadyExists
		}
	}
	m[key] = value
	return m, nil
}

/*---------------------------------------------------*/
func (m MyMaps) Delete(key int, value string) (MyMaps, error) {
	newM5 := MyMaps{}
	for i, _ := range m {
		if i == key && m[i] == value {
			continue
		}
		newM5[i] = m[i]
	}
	if AreMapsSame(m, newM5) {
		return m, Error_NotFound
	}
	return newM5, nil
}

/*---------------------------------------------------*/
func (m MyMaps) Update(key int, value string) MyMaps {
	m[key] = value
	return m
}

/*---------------------------------------------------*/
