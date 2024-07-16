# Slack Bot YOB Calculator

Welcome to the Slack Bot YOB (Year of Birth) Calculator project! This is a simple Slack bot implemented in Go using the `slacker` library. The bot calculates your age based on the year of birth provided in a Slack command.

## Getting Started

### Prerequisites

- Go installed on your machine
- Slack App and Bot tokens
- slack App Token := https://api.slack.com/apps/A07CYPFDG3T/general
- slack Bot token :=https://app.slack.com/app-settings/T07CKD0RUDS/A07CYPFDG3T/socket-mode

### Configuration

Set the following environment variables with your Slack Bot Token and Slack App Token:

```bash
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
```
### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/uday778/Golang-handbook.git
    ```
2. Navigate to the project directory:
    ```bash
    cd p4_slack-age-bot
    ```
3. Run the application:
    ```bash
    go run main.go
    ```
The bot will start listening for Slack commands.

## command 
```bash
@Age Bot my yob is 1997
```

### The bot will reply with your calculated age.

## Command Events

The application prints command events to the console. These events include details about each command, such as timestamp, command name, parameters, and event type.

## Acknowledgments

- [slacker](https://github.com/shomali11/slacker): A Slack client library for Go.
- [AkhilSharma90/GO-Slackbot-Calculates-Age](https://github.com/AkhilSharma90/GO-Slackbot-Calculates-Age): Acknowledging the source for inspiration.