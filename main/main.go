package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	wd, _ := os.Getwd()

	script := `
import { hello } from 'mymod';

window.hello = hello;

`

	importResolver := api.Plugin{
		Name: "hugo-import-resolver",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `.*`},
				func(args api.OnResolveArgs) (api.OnResolveResult, error) {
					if args.Path == "mymod" {
						return api.OnResolveResult{
							Path: filepath.Join(wd, "..", "module/src/hello.js"),
						}, nil
					}
					return api.OnResolveResult{}, nil

				})
		},
	}

	opts := api.BuildOptions{
		Bundle: true,

		Stdin: &api.StdinOptions{
			Contents:   script,
			ResolveDir: wd,
			Loader:     api.LoaderJS,
		},

		Plugins: []api.Plugin{importResolver},
	}

	res := api.Build(opts)

	if res.Errors != nil {
		log.Fatal(res.Errors[0].Text)
	}
}
