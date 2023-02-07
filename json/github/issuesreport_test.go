package github

import (
	"html/template"
	"log"
	"os"
	"testing"
)

func TestTempl(t *testing.T) {
	//report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	//if err != nil {
	//	log.Fatal(err)
	//}

	issue := []string{"repo:golang/go", "is:open", "json", "decoder"}
	//sample 1
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	result, err := SearchIssues(issue)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	//sample 2
	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	 <th>#</th>
	 <th>State</th>
	 <th>User</th>
	 <th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
	 <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	 <td>{{.State}}</td>
	 <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	 <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))
	result, err = SearchIssues(issue)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
