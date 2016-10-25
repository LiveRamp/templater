package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var (
	source = flag.String("template", "", "Template to render")
	data   = flag.String("data", "", "Data file to use as template context")
)

func main() {
	flag.Parse()

	if *source == "" || *data == "" {
		fatal("please pass a template and data")
	}

	var deserialized map[string]interface{}
	if fp, err := os.Open(*data); err != nil {
		fatal(fmt.Sprintf("failed to open data file: %s", err))
	} else if err = json.NewDecoder(fp).Decode(&deserialized); err != nil {
		fatal(fmt.Sprintf("failed to deserialize data file: %s", err))
	} else if t, err := template.New("").ParseFiles(*source); err != nil {
		fatal(fmt.Sprintf("failed to parse template: %s", err))
	} else if err = t.ExecuteTemplate(os.Stdout, filepath.Base(*source), deserialized); err != nil {
		fatal(fmt.Sprintf("failed to render template: %s", err))
	}
}

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, "FATAL: %s\n", msg)
	os.Exit(1)
}
