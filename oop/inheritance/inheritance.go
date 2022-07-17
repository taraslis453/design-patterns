package inheritance

import "fmt"

type author struct {
	name  string
	email string
}

func (a *author) getAuthorDetails() string {
	return a.name + " " + a.email
}

type blogPost struct {
	author
	title string
	body  string
}

func (b *blogPost) getBlogPostDetailsInMarkdown() string {
	return "Title: " + b.title + "\n" + "Body: " + b.body + "\n" + "Author: " + b.getAuthorDetails()
}

// go doesn't have a concept of inheritance but it can be implemented using composition
func RunInheritance() {
	b := blogPost{
		author: author{
			name:  "John Doe",
			email: "johndoe@email.com",
		},
		title: "My first blog post",
		body:  "This is my first blog post",
	}
	fmt.Println(b.getBlogPostDetailsInMarkdown())
}
