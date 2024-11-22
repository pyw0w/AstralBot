package pluginBase

import (
	"fmt"

	"github.com/pyw0w/AstralBot/logger"
)

type Plugin interface {
	Init() error
	Run() error
}

var plugins []Plugin

func RegisterPlugin(p Plugin) {
	plugins = append(plugins, p)
}

func InitPlugins() error {
	for _, p := range plugins {
		logger.Info("pluginBase", "Initializing plugin: "+getPluginName(p))
		if err := p.Init(); err != nil {
			return err
		}
	}
	return nil
}

func RunPlugins() error {
	for _, p := range plugins {
		logger.Info("pluginBase", "Running plugin: "+getPluginName(p))
		if err := p.Run(); err != nil {
			return err
		}
	}
	return nil
}

func Start() error {
	if err := InitPlugins(); err != nil {
		return err
	}

	if err := RunPlugins(); err != nil {
		return err
	}

	logger.Info("pluginBase", "All plugins started successfully")
	return nil
}

func getPluginName(p Plugin) string {
	return fmt.Sprintf("%T", p)
}
