package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

type TemplateData struct {
	Countries []Country
}

func main() {
	http.HandleFunc("/", countriesHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		http.Error(w, "Failed to fetch countries data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		http.Error(w, "Failed to parse countries data", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("countries").Parse(htmlTemplate))
	data := TemplateData{Countries: countries}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Countries Table</title>
  <style>
    table {
      width: 100%;
      border-collapse: collapse;
    }
    table, th, td {
      border: 1px solid black;
    }
    th, td {
      padding: 8px;
      text-align: left;
    }
    th {
      background-color: #f2f2f2;
    }
    img {
      width: 32px;
      height: 32px;
    }
  </style>
</head>
<body>
  <h2>Countries Table</h2>
  <table>
    <thead>
      <tr>
        <th>Flag</th>
        <th>Name</th>
        <th>Dial Code</th>
        <th>ISO Code</th>
      </tr>
    </thead>
    <tbody>
      {{range .Countries}}
      <tr>
        <td><img src="{{.Flag}}" alt="{{.Name}} flag"></td>
        <td>{{.Name}}</td>
        <td>{{.DialCode}}</td>
        <td>{{.IsoCode}}</td>
      </tr>
      {{end}}
    </tbody>
  </table>
</body>
</html>
`
