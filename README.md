![DALL·E 2022-08-04 20 18 34 - very cute happy blue cartoon gopher send a message with computer(1)](https://user-images.githubusercontent.com/109267960/182912478-cc4e4910-5524-4067-993d-1696db25358c.png)

## Thanks to AI'S
This package get helps from AI's Github Copilot, OpenAI and DALL·E

## Installation

As a library

```shell
go get github.com/utfd/slack-alerts
```

or if you want to use it as a bin command

go >= 1.17
```shell
go install github.com/utfd/slack-alerts@latest
```

go < 1.17
```shell
go get github.com/utfd/slack-alerts
```
### Setting webhooks
At your .env file set SLACK_WEBHOOK_URL like bellow.
```.env
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXXXX
```

## Usage

To send a message to Slack, simply call `SlackAlert::message()` and pass it any message you want.

```go
slackAlerts.Message("You have a new subscriber to the Foo newsletter!")
```


## Formatting

### Markdown
You can format your messages with Slack's markup. Learn how [in the Slack API docs](https://slack.com/help/articles/202288908-Format-your-messages).

```go
slackAlerts.Message("A message *with some bold statements* and _some italicized text_.")
```

Links are formatted differently in Slack than the classic markdown structure.

```go
slackAlerts.Message("<https://google.com|This is a link to our homepage>")
```

### Emoji's

You can use the same emoji codes as in Slack. This means custom emoji's are also supported.
```go
slackAlerts.Message(":smile: :custom-code:")
```

### Mentioning

You can use mentions to notify users and groups. Learn how [in the Slack API docs](https://api.slack.com/reference/surfaces/formatting#mentioning-users).
```go
slackAlerts.Message("A message that notifies <@username> and everyone else who is <!here>")
```

## Testing

```bash
go test
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
