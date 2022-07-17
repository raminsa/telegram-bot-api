# Telegram Bot API

The Telegram Bot API provides an HTTP API for creating Telegram Bots.

If you've got any questions about bots or would like to report an issue with your bot, kindly contact us at @BotSupport in Telegram.

Please note that only global Bot API issues that affect all bots are suitable for this repository.

To learn how to use it, please see our [examples](https://github.com/Raminsa/Telegram_API/tree/main/examples).

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Example](#example)
- [Custom Client](#custom-client)
- [Debug](#debug)

<a name="installation"></a>
## Installation
`go get -u github.com/Raminsa/Telegram_API`.

<a name="documentation"></a>
## Documentation
See [Bots: An introduction for developers](https://core.telegram.org/bots) for a brief description of Telegram Bots and their features.

See the [Telegram Bot API documentation](https://core.telegram.org/bots/api) for a description of the Bot API interface and a complete list of available classes, methods and updates.

See the [Telegram Bot API server build instructions generator](https://tdlib.github.io/telegram-bot-api/build.html) for detailed instructions on how to build the Telegram Bot API server.

Subscribe to [@BotNews](https://t.me/botnews) to be the first to know about the latest updates and join the discussion in [@BotTalk](https://t.me/bottalk).


<a name="example"></a>
## Example

use get update method (simple):
```go
package main

import (
	"fmt"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
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

	"github.com/Raminsa/Telegram_API/telegram"
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

	"github.com/Raminsa/Telegram_API/telegram"
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


<a name="custom-client"></a>
## Custom Client
use client with custom options:
```go
package main

import (
	"fmt"
	"log"

	"github.com/Raminsa/Telegram_API/telegram"
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

	"github.com/Raminsa/Telegram_API/telegram"
)

func main() {
	tg, err := telegram.NewWithBaseUrl("BotToken","baseUrl")
	if err != nil {
		log.Fatal(err)
	}

	tg.Bot.Debug = true
	me, err := tg.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("botID:", me.ID, "botUsername:", me.UserName)

	err = tg.WriteLoggerFile("fileName")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tg.GetLoggerFile())
}
```
