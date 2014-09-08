package main

import (
        "bufio"
        "bytes"
        "log"
        "os"
        "strings"
        "github.com/codegangsta/cli"
        "github.com/usami/slackcli/slack"
)

func main() {
        app := cli.NewApp()
        app.Name = "slackcli"
        app.Usage = "post messeges to Slack from stdin"

        app.Flags = []cli.Flag {
                cli.StringFlag{"channel, c", os.Getenv("SLACK_CHANNEL"), "specify channel"},
                cli.StringFlag{"username, u", os.Getenv("SLACK_USERNAME"), "specify username"},
                cli.StringFlag{"icon-url", os.Getenv("SLACK_ICON_URL"), "set icon with url"},
                cli.StringFlag{"icon-emoji", os.Getenv("SLACK_ICON_EMOJI"), "set emoji icon"},
        }

        app.Action = func(c *cli.Context) {
                subdomain := os.Getenv("SLACK_SUBDOMAIN")
                token := os.Getenv("SLACK_TOKEN")

                if len(subdomain) == 0 || len(token) == 0 {
                        log.Fatal("You must set the environment variables: SLACK_SUBDOMAIN and SLACK_TOKEN.")
                }

                var buf bytes.Buffer
                scanner := bufio.NewScanner(os.Stdin)

                for scanner.Scan() {
                        buf.WriteString(scanner.Text())
                        buf.WriteString("\n")
                }
                if scanner.Err() != nil {
                        panic(scanner.Err())
                }

                payload := slack.BuildPayload(c.String("channel"), c.String("username"),
                                              c.String("icon-url"), c.String("icon-emoji"),
                                              strings.Trim(buf.String(), "\n"))
                slack.Post(subdomain, token, payload)
        }
        app.Run(os.Args)
}
