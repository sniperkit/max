/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/sniperkit/snk.fork.max/cmd"
)

var version = "master"

func main() {
	cmd.Version = version
	cmd.Execute()
}
