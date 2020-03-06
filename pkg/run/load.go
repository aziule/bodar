package run

import (
	"fmt"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/behaviour/http"
	"github.com/aziule/bodar/pkg/config"
)

type Loader interface {
	Load(cfg *config.UserConfig) error
}

// ConfigLoader is a struct used to load behaviours.
type ConfigLoader struct {
	runner *Runner
}

// NewLoader creates a new ConfigLoader.
func NewLoader(runner *Runner) *ConfigLoader {
	return &ConfigLoader{
		runner: runner,
	}
}

func (l *ConfigLoader) Load(cfg *config.UserConfig) error {
	l.loadDefaultBehaviours()

	for _, bCfg := range cfg.Behaviours {
		_, ok := l.runner.available[bCfg.Name]
		if !ok {
			return fmt.Errorf(`behaviour "%s" not found`, bCfg.Name)
		}

		l.runner.Use(bCfg.Name, bCfg.Params)
	}

	//httpServerCfgs := make(map[string]behaviour.Config)
	//
	//for name, serverConfig := range cfg.Servers.Http {
	//	httpServerCfgs[name] = behaviour.Config{
	//		"read_timeout":        serverConfig.ReadTimeout,
	//		"read_header_timeout": serverConfig.ReadHeaderTimeout,
	//		"write_timeout":       serverConfig.WriteTimeout,
	//		"idle_timeout":        serverConfig.IdleTimeout,
	//		"max_header_bytes":    serverConfig.MaxHeaderBytes,
	//	}
	//}
	//
	//for _, behaviourConfig := range cfg.Behaviours.Http {
	//	serverCfg, ok := httpServerCfgs[behaviourConfig.Server]
	//	if !ok {
	//		return fmt.Errorf(`server "%s" not defined for behaviour "%s"`, behaviourConfig.Server, behaviourConfig.Type)
	//	}
	//
	//	srv, err := http.NewDefaultServer(serverCfg)
	//	if err != nil {
	//		return fmt.Errorf(`could not create server for behaviour "%s": %v`, behaviourConfig.Type, err)
	//	}
	//
	//	_, ok = l.runner.available[behaviourConfig.Type]
	//	if !ok {
	//		return fmt.Errorf(`behaviour type "%s" not found`, behaviourConfig.Type)
	//	}
	//
	//	bCfg := behaviour.Config{
	//		"server": srv,
	//		"port":   behaviourConfig.Port,
	//	}
	//	for k, v := range behaviourConfig.Params {
	//		bCfg[k] = v
	//	}
	//
	//	l.runner.Use(behaviourConfig.Type, bCfg)
	//	fmt.Println(l.runner)
	//	//srv := http.NewDefaultServer(http.DefaultServerConfig{})
	//	//
	//	//r := (&run.Runner{}).WithDefaultStrategies()
	//	//r.Use(http.EmptyBodyBehaviourName, map[string]interface{}{
	//	//	"port":   8081,
	//	//	"server": srv,
	//	//})
	//	//r.Use(http.StatusCodeBehaviourName, map[string]interface{}{
	//	//	"port":        8082,
	//	//	"server":      srv,
	//	//	"status_code": 404,
	//	//})
	//}

	return nil
}

func (l *ConfigLoader) LoadCustomBehaviour(name string, factoryFunc behaviour.FactoryFunc) *ConfigLoader {
	l.runner.available[name] = factoryFunc
	return l
}

func (l *ConfigLoader) loadDefaultBehaviours() *ConfigLoader {
	l.runner.available[http.EmptyBodyBehaviourName] = http.NewEmptyBodyBehaviour
	l.runner.available[http.StatusCodeBehaviourName] = http.NewStatusCodeBehaviour
	return l
}
