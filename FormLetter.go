package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var newEvents = []string{"Aug 15 : Independence Day India", "October 2 : Gandhi Birthday", "Dec 31 : New year"}

	// Template stuct for injecing data in form
	type FunraiseDetails struct {
		HonorificName     string
		Donation          bool
		WasThere          bool
		NewEventForFuture []string
	}

	const lettertemp = `
  Dear {{.HonorificName}},
  {{if .WasThere}}
  It was a pleasure to see you attend the Fundraiser.{{else}}
  It is a shame you couldn't make it to the Fundraiser.{{end}}
  {{if .Donation}}Thank you for donation.
  It is Greatly appreciated.
  {{end}}
  We would be happy to have you for our upcoming events:{{range $element := .NewEventForFuture}} {{$element}}{{end}}

  Regards,
  Ishan KAushik
`

	var guests = []FunraiseDetails{
		{"Mr Jay Kumar", true, true, newEvents},
		{"Miss Daisy", true, false, newEvents},
		{"Mr Lambbottom", false, false, newEvents},
	}

	// Createandparse trmplate
	t := template.Must(template.New("letter").Parse(lettertemp))

	// Iterte foreach guest
	for _, eachGuest := range guests {
		err := t.Execute(os.Stdout, eachGuest)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
