package main

import (
  "os"
  "github.com/noirbizarre/gonja"
  "github.com/adrg/frontmatter"
  "github.com/urfave/cli/v2"
  "github.com/Joker/hpp"
  "github.com/brothertoad/btu"
)

func main() {
  app := &cli.App {
    Name: "gonjademo",
    Usage: "a little program to demonstrate gonja templates",
    Action: doDemo,
    Flags: []cli.Flag {
      &cli.StringFlag { Name: "template", Usage: "template file to parse", Required: true, },
      &cli.StringFlag { Name: "output", Usage: "output file", Required: true, },
    },
  }
  app.Run(os.Args)
}

func doDemo(c *cli.Context) error {
  // Extract the frontmatter from the template, then parse what's left.
  fm := make(map[string]interface{})
  reader := btu.OpenFile(c.String("template"))
  defer reader.Close()
  rest, fmerr := frontmatter.Parse(reader, &fm)
  btu.CheckError(fmerr)
  gonja.DefaultEnv.KeepTrailingNewline = true
  tpl := gonja.Must(gonja.FromBytes(rest))
  out, err := tpl.Execute(fm)
  btu.CheckError(err)
  err = os.WriteFile(c.String("output"), []byte(hpp.PrPrint(out)), 0644)
  btu.CheckError(err)
  return nil
}
