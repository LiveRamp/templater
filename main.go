package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"text/template"

	log "github.com/Sirupsen/logrus"
)

var (
	source = flag.String("template", "", "Template to render")
	data   = flag.String("data", "", "Data file to use as template context")
)

func main() {
	flag.Parse()

	if *source == "" || *data == "" {
		log.Fatal("please pass a template and data")
	}

	var deserialized map[string]interface{}
	if fp, err := os.Open(*data); err != nil {
		log.WithField("err", err).Fatal("failed to open data file")
	} else if err = json.NewDecoder(fp).Decode(&deserialized); err != nil {
		log.WithField("err", err).Fatal("failed to deserialize data file")
	} else if t, err := template.New("").ParseFiles(*source); err != nil {
		log.WithField("err", err).Fatal("failed to parse template")
	} else if err = t.ExecuteTemplate(os.Stdout, filepath.Base(*source), deserialized); err != nil {
		log.WithField("err", err).Fatal("failed to render template")
	}
}
