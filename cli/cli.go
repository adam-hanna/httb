package cli

import (
	"strconv"

	"github.com/urfave/cli"
)

// FlagsStruct - holds command line args
type FlagsStruct struct {
	URL      string
	Headers  string
	Cookies  string
	HTTPCode int
	Method   string
}

// StartCLI - gathers command line args
func StartCLI(cliFlags *FlagsStruct, args []string) error {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		var err error

		URL := ctx.GlobalString("url")
		Headers := ctx.GlobalString("headers")
		Cookies := ctx.GlobalString("cookies")
		HTTPCode := ctx.GlobalString("http-code")
		Method := ctx.GlobalString("http-method")

		// build the cli struct to send back to main
		cliFlags.URL = URL
		cliFlags.Headers = Headers
		cliFlags.Cookies = Cookies
		cliFlags.HTTPCode, err = strconv.Atoi(HTTPCode)
		cliFlags.Method = Method

		return err
	}
	app.Authors = []cli.Author{
		{
			Email: "ahanna@alumni.mines.edu",
			Name:  "Adam Hanna",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "URL to run benchmark against (include http://)",
			Value: "http://localhost:8080",
		},
		cli.StringFlag{
			Name:  "headers, z",
			Usage: "Headers to include (e.g. key1=val1;key2=val2)",
			Value: "",
		},
		cli.StringFlag{
			Name:  "cookies, c",
			Usage: "Cookies to include (e.g. cookie1=value1;cookie2=value2)",
			Value: "",
		},
		cli.StringFlag{
			Name:  "http-code, i",
			Usage: "Expected http status code (e.g. 200)",
			Value: "200",
		},
		cli.StringFlag{
			Name:  "http-method, m",
			Usage: "HTTP method to use (e.g. GET)",
			Value: "GET",
		},
	}
	app.Name = "httb"
	app.Usage = "An http benchmarking tool"
	app.Version = "0.0.1"
	return app.Run(args)
}
