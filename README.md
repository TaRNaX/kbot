#### Telegram bot

Telegram bot on golang using ![cobra-cli](https://github.com/spf13/cobra-cli)

Installation instructions:
- Fork or clone repo to your github account
- Open project in GitHub codespaces
- To add more functions to bot install cobra-cli tool by `go install github.com/spf13/cobra-cli`
- Create telegram bot and generate token by 'BotFather'
- Export token to environment variable `export TELE_TOKEN=<our token>`
- Start bot `./kbot start/`
- Open bot in telegram

Link to access telegram bot:
- t.me/mtarnax_bot

Telegram bot commands:

- `/start hello`  welcome message with current version of app
- `/start joke`   returns a random joke
- `/start fact`   returns a random interesting fact
