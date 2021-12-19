package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	DEBUG        bool
	TCache       map[string]*template.Template
	InProduction bool
	InfoLog      *log.Logger
	Session      *scs.SessionManager
}
