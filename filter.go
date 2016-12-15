package main

import "github.com/urfave/cli"

func filterConfigs(configs []Config, context *cli.Context) []Config {
	if !context.GlobalBool("passwords") {
		for i := 0; i < len(configs); i++ {
			configs[i].Password = ""
		}
	}
	return configs
}
