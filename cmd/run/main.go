package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/aziule/bodar/pkg/config"
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

	data, err := ioutil.ReadFile(".bodar.yml")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	cfg, err := config.Parse(data)
	if err != nil {
		log.Fatalf("could not parse config: %v", err)
	}

	runner := &run.Runner{}
	loader := run.NewLoader(runner)
	err = loader.Load(cfg)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	err = runner.Run(context.Background())
	if err != nil {
		applog.Fatalf("error starting the runner: %v")
	}
}
