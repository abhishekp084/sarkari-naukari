// models.go
package main

type Content struct {
	Heading string
	Text    string
}

type Job struct {
	Title       string
	Description string
	Salary      string
}

type PageData struct {
	Title       string
	ContentList []Content
	Jobs        []Job
}
