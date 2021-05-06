package core

import (
	"bytes"
	"text/template"
)

func ParseConnectionURL(cfg *Config) (string, error) {
	f := "host={{.Host}} port={{.Port}} user={{.Username}} password={{.Password}} dbname={{.DbName}} sslmode=disable"
	tmpl, err := template.New("database").Parse(f)
	if err != nil {
		return "", err
	}

	return fillConnectionURL(cfg, tmpl)
}

func ParseMigrationConnectionURL(cfg *Config) (string, error) {
	f := "{{.Driver}}://{{.Username}}:{{.Password}}@{{.Host}}:{{.Port}}/{{.DbName}}?sslmode=disable"
	tmpl, err := template.New("migration").Parse(f)
	if err != nil {
		return "", err
	}

	return fillConnectionURL(cfg, tmpl)
}

func fillConnectionURL(cfg *Config, tmpl *template.Template) (string, error) {
	var conn bytes.Buffer
	err := tmpl.Execute(&conn, cfg.Database)
	if err != nil {
		return "", err
	}

	return conn.String(), nil
}
