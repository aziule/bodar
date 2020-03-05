package main

import (
	"context"
	"log"
	"os"

	"github.com/aziule/bodar/pkg/behaviour/http"
	applog "github.com/aziule/bodar/pkg/log"
	"github.com/aziule/bodar/pkg/run"
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

	srv := http.NewDefaultServer(http.DefaultServerConfig{})

	r := (&run.Runner{}).WithDefaultStrategies()
	r.Use(http.EmptyBodyBehaviourName, map[string]interface{}{
		"port":   8081,
		"server": srv,
	})

	err = r.Run(context.Background())
	if err != nil {
		applog.Fatalf("error starting the runner: %v")
	}
}
