package blogposts

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	desctiptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	bodySeparator        = "Body: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}
type StubFailinfFS struct {
}

func (s StubFailinfFS) Open(name string) (fs.File, error) {
	return nil, errors.New("danm, i failed")
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {

		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}
func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(desctiptionSeparator)
	tags := strings.Split(readLine(tagsSeparator), ", ")

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "/n")
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body[6:],
	}, nil
}
