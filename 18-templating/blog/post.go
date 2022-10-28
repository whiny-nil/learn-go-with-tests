package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	return Post{
		Title:       readMetaLine(scanner, titleSeparator),
		Description: readMetaLine(scanner, descriptionSeparator),
		Tags:        readTags(scanner),
		Body:        readBody(scanner),
	}, nil
}

func readMetaLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}

func readTags(scanner *bufio.Scanner) []string {
	tags := strings.Split(readMetaLine(scanner, tagsSeparator), ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}
	return tags
}

func readBody(scanner *bufio.Scanner) string {
	// throw away '---' separator
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
