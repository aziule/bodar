package main

import (
	"log"
	"os"

	"github.com/aziule/bodar/internal/app/bodar"
	applog "github.com/aziule/bodar/internal/app/log"
	"github.com/aziule/bodar/pkg/strategy/http"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.Out = os.Stdout
	l.SetLevel(logrus.InfoLevel)
	err := applog.Setup(l)
	if err != nil {
		log.Fatalf("could not setup logger: %v", err)
	}

	r := &bodar.Registry{}

	srv := http.NewDefaultServer(http.DefaultServerConfig{})

	r.Use(http.EmptyBodyStrategyName, map[string]interface{}{
		"port":   8081,
		"server": srv,
	})
	log.Fatalf("error running registry: %v", r.Run())
}
