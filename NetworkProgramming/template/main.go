package main

import (
	"html/template"
	"fmt"
	"bytes"
	"net/http"
)

var (
	bufStrings string
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}
The age is {{.Age}}

{{range .Emails}}
	An email is {{.}}
{{end}}

{{with.Jobs}}

	{{range .}}
		An employer is {{.Employer}}
		and the role is {{.Role}}
	{{end}}

{{end}}
`

func main() {

	job1 := Job{Employer: "Monash", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of He"}

	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"jan@dsfsdf.com", "jan.dsfdf@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)

	if err != nil {
		panic(err)
	}
	strings := bytes.NewBuffer(nil)
	
	err = t.Execute(strings, person)
	fmt.Println(strings)

	var buf [512]byte
	
	for {
		n,err := strings.Read(buf[0:])
		if err != nil {
			break
		}
		bufStrings += string(buf[0:n])
	}

	http.HandleFunc("/", print)
	err = http.ListenAndServe(":8000", nil) 
	if err != nil {
		panic(err)
	}
}

func print(writer http.ResponseWriter, req *http.Request) {
writer.Write([]byte(bufStrings))

}
