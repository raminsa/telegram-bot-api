# Telegram Bot API

The Telegram Bot API provides an HTTP API for creating Telegram Bots.

If you've got any questions about bots or would like to report an issue with your bot, kindly contact us at @BotSupport in Telegram.

Please note that only global Bot API issues that affect all bots are suitable for this repository.

To learn how to use it, please see our [examples](https://github.com/raminsa/telegram-bot-api/tree/main/examples).

Bot API 7.2 recent changes [March 31, 2024](https://core.telegram.org/bots/api#march-31-2024).

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Example](#example)
- [Custom Client](#custom-client)
- [Debug](#debug)

<a name="installation"></a>
## Installation
`go get github.com/raminsa/telegram-bot-api`.

<a name="documentation"></a>
## Documentation
See [Bots: An introduction for developers](https://core.telegram.org/bots) for a brief description of Telegram Bots and their features.

See the [Telegram Bot API documentation](https://core.telegram.org/bots/api) for a description of the Bot API interface and a complete list of available classes, methods and updates.

See the [Telegram Bot API server build instruction generator](https://tdlib.github.io/telegram-bot-api/build.html) for detailed instructions on how to build the Telegram Bot API server.

Subscribe to [@BotNews](https://t.me/botnews) to be the first to know about the latest updates and join the discussion in [@BotTalk](https://t.me/bottalk).


<a name="example"></a>
## Example

use get update method (simple):
```go
package main

import (
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	getUpdates := tg.NewGetUpdates()
	getUpdates.Offset = 1
	getUpdates.Timeout = 60
	getUpdates.Limit = 100
	updates := tg.GetUpdatesChan(getUpdates)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		fmt.Println(update.UpdateID, getUpdates.Offset)
		if update.UpdateID >= getUpdates.Offset {
			getUpdates.Offset = update.UpdateID + 1
		}
		if update.Message != nil {
			fmt.Println(update.Message.Text)
		}
	}
}
```

use webhook method (http):
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")
	err := http.ListenAndServe("BotPortNumber", http.HandlerFunc(handleWebhook))
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := telegram.HandleUpdate(r)
	if err != nil {
		telegram.HandleUpdateError(w, err)
		return
	}

	if update.Message != nil {
		fmt.Println(update.Message.Text)
	}
}
```

use webhook method (https):
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")
	err := http.ListenAndServeTLS("BotPortNumber", "BotCertFile", "BotKeyFile", http.HandlerFunc(handleWebhook))
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := telegram.HandleUpdate(r)
	if err != nil {
		telegram.HandleUpdateError(w, err)
		return
	}

	if update.Message != nil {
		fmt.Println(update.Message.Text)
	}
}
```

to generate your cert file use this. See [self-signed](https://core.telegram.org/bots/self-signed) guide for details.:

    openssl req -newkey rsa:2048 -sha256 -nodes -keyout <file.key> -x509 -days 36500 -out <file.pem> -subj "/C=US/ST=New York/L=Brooklyn/O=Example Brooklyn Company/CN=<server_address>"


use a channel to handle all requests (Avoid ReadTimeoutExpired error), http.ListenAndServe() supports concurrency:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/raminsa/telegram-bot-api/telegram"
	"github.com/raminsa/telegram-bot-api/types"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")
	updates := listenForWebhook(100)
	go http.ListenAndServeTLS("BotPortNumber", "BotCertFile", "BotKeyFile", nil)
	for update := range updates {
		if update.Message != nil {
			fmt.Println(update.Message.Text)
		}
	}
}

func listenForWebhook(maxWebhookConnections int) types.UpdatesChannel {
	ch := make(chan types.Update, maxWebhookConnections)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if update, err := telegram.HandleUpdate(r); err != nil {
			return
		} else {
			ch <- update
		}
	})

	return ch
}
```

<a name="custom-client"></a>
## Custom Client
use client with custom options:
```go
package main

import (
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	client := telegram.Client()
	client.BaseUrl = "baseUrl"
	client.Proxy = "proxy"
	client.ForceV4 = true
	client.DisableSSLVerify = true
	client.ForceAttemptHTTP2 = true
	tg, err := telegram.NewWithCustomClient("BotToken", &client)
	if err != nil {
		log.Fatal(err)
	}

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)
}
```

<a name="debug"></a>
## Debug

use debug option and write to local file:
```go
package main

import (
	"fmt"
	"log"

	"github.com/raminsa/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.NewWithBaseUrl("BotToken", "baseUrl")
	if err != nil {
		log.Fatal(err)
	}

	//active debug mode
	tg.Bot.Debug = true

	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)

	//access debug log
	//method 1: write to console
	fmt.Println(tg.GetLoggerFile())

	//method 2: write to file
	err = tg.WriteLoggerFile("fileName")
	if err != nil {
		log.Fatal(err)
	}
}
```
