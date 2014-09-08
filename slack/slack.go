package slack

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        )

type Payload struct {
        Content
        IconURL string `json:"icon_url,omitempty"`
        IconEmoji string `json:"icon_emoji,omitempty"`
}

type Content struct {
        Channel string `json:"channel"`
        Username string `json:"username"`
        Text string `json:"text"`
}

func BuildPayload(channel, username, icon_url, icon_emoji, text string) *bytes.Reader {
        if len(icon_url) != 0 {
                icon_emoji = ""
        }

        content := Content{channel, username, text}
        payload := Payload{content, icon_url, icon_emoji}

        b, err := json.Marshal(payload)
        if err != nil {
                panic(err)
        }
        return bytes.NewReader(b)
}


func Post(subdomain, token string, payload *bytes.Reader) {
        url := fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?token=%s", subdomain, token)
        resp, err := http.Post(url, "application/json", payload)
        defer resp.Body.Close()

        if err != nil {
                panic(err)
        }

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err)
        }

        fmt.Printf("%s\n", string(body))
}
