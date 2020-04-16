<h1 align="center">
  <br>
  <img src="logo-hellper.png" alt="Hellper - Your best friend in times of crisis" width="222"><br>
  Hellper - Your best friend in times of crisis<br>
</h1>

<div align="center">Hellper bot aims to orchestrate the process and resolution of incidents, reducing the time spent with manual tasks and ensuring that the necessary steps are fulfilled in the right order. Also, it facilitates the measurement of impact and response rate through <a href="#metrics">metrics</a>.
  <p>A chance to help explore and develop a bot written in Go, integrated with multiple external platforms and tools.</p>
  <p>Help us expand incident processes’ and understand the needs of other companies that may benefit from Hellper bot.</p>
  <p>You’re just one PR away from joining the developing team of Hellper! <a href="#contributing">Contribute</a></p>
</div>

<p align="center">
  <a href="https://circleci.com/gh/ResultadosDigitais/hellper"><img alt="CircleCI" src="https://circleci.com/gh/ResultadosDigitais/hellper.svg?style=svg&circle-token=66f54e118b9316ddfb9357299268c42dc772df04"></a>
  <a href="https://dependabot.com"><img alt="Dependabot Status" src="https://api.dependabot.com/badges/status?host=github&repo=ResultadosDigitais/hellper&identifier=254472121"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/PRs-welcome-informational.svg" alt="PRs welcome!" /></a>
  <a href="#license"><img alt="License" src="https://img.shields.io/badge/license-MIT-informational"></a>
</p>

---

## Contents
1. [Getting Started](#getting-started)
   * [Prerequisites](#prerequisites)
   * [Installing](#installing)
2. [Running the Tests](#running-the-tests)
3. [Running the Application](#running-the-application)
4. [Deployment](#deployment)
5. [Optional Setup](#optional-setup)
   * [Ngrok (To receive events from Slack)](#ngrok-to-receive-events-from-slack)
   * [Setup Golang](#setup-golang)
   * [Setup Database](#setup-database)
6. [How to use](#how-to-use)
   * [Commands](#commands)
   * [Metrics](#metrics)
7. [Contributing](#contributing)
8. [Code of Conduct](#code-of-conduct)
9. [Need help?](#need-help)
10. [License](#license)


## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites
1. [Docker Compose](https://github.com/docker/compose/releases)
2. [Slack Account](https://slack.com/)
3. [G Suite Account](https://gsuite.google.com/)

### Installing
1. Clone this repo
```
git clone git@github.com:ResultadosDigitais/hellper.git
```

2. [Configure Slack](/docs/CONFIGURING-SLACK.md)
3. [Configure Google](/docs/CONFIGURING-GOOGLE.md)
4. Make a copy from configuration example
```
cp development.env.example development.env
```

#### Variables explanation
| Variable | Explanation | Default value |
| --- | --- | --- |
|**HELLPER_LANGUAGE**|Hellper languague|`pt-br`|
|**HELLPER_BIND_ADDRESS**|Hellper local bind address|`:8080`|
|**HELLPER_DATABASE**|Database provider (supported values: postgres)| `postgres` |
|**HELLPER_DSN**|Your Data Source Name| --- |
|**HELLPER_ENVIRONMENT**|Current environment (supported values: production, staging)| --- |
|**HELLPER_GOOGLE_DRIVE_CREDENTIALS**|[Google Drive Credentials](/docs/CONFIGURING-GOOGLE.md#Get-a-Client-ID-and-Client-Secret)| --- |
|**HELLPER_GOOGLE_DRIVE_FILE_ID**|[Google Drive FileId](/docs/CONFIGURING-GOOGLE.md#Template-Post-mortem) to your post-mortem template| --- |
|**HELLPER_GOOGLE_DRIVE_TOKEN**|[Google Drive Token](/docs/CONFIGURING-GOOGLE.md#Signing-in-to-the-application)| --- |
|**HELLPER_MATRIX_HOST**|[Matrix](https://github.com/ResultadosDigitais/matrix) URL host| --- |
|**HELLPER_PRODUCT_CHANNEL_ID**|The Product channel id used to notify new incidents| --- |
|**HELLPER_SUPPORT_TEAM**|Support team identifier to notify| --- |
|**HELLPER_REMINDER_STATUS_SECONDS**|Contains the time for the stat reminder to be triggered in open incidents, by default the time is 2 hours if there is no variable| `7200` |
|**HELLPER_OAUTH_TOKEN**|[Slack token](/docs/CONFIGURING-SLACK.md#User-Token-Scopes) to exeucte bot user actions| --- |
|**HELLPER_VERIFICATION_TOKEN**|[Slack token](/docs/CONFIGURING-SLACK.md#User-Token-Scopes) to verify external requests| --- |
|**HELLPER_NOTIFY_ON_RESOLVE**|Notify the main channel when resolve the incident| `true` |
|**HELLPER_NOTIFY_ON_CLOSE**|Notify the main channel when close the incident| `true` |
|**FILE_STORAGE**|Hellper file storage for postmortem document| `google_drive` |

## Running the Tests
1. `make test`

## Running the application
2. `make run`
___

## Deployment
_Coming soon_

## Optional Setup

### Ngrok (To receive events from Slack)
- Download Ngrok https://ngrok.com/ and create account
- `sudo cp ngrok /usr/local/bin`
- `ngrok http 8080`
- Copy your public address. You'll need this to [Configure Slack API](/docs/CONFIGURING-SLACK.md)

### Golang
- Install [golang](https://golang.org/doc/install)

OR

1. Install [gvm](https://github.com/moovweb/gvm)
2. Follow gvm post install instructions
3. Install go 1.13 as default

### Database
```shell
psql $HELLPER_DSN -f "./internal/model/sql/postgres/schema/hellper.sql"
```

## How to use

### Commands
After [Configuring Slack](/docs/CONFIGURING-SLACK.md) you can use the commands created. The commands are as it follows:

| Command  | Request URL | Short Description |
| - | - | - |
|`/hellper_incident`|https://yourhost.publicaddress.com/open|_Starts Incident_|
|`/hellper_status`|https://yourhost.publicaddress.com/status|_Show all pinned messages_|
|`/hellper_close`|https://yourhost.publicaddress.com/close|_Closes Incident_|
|`/hellper_resolve`|https://yourhost.publicaddress.com/resolve|_Resolves Incident_|
|`/hellper_cancel`|https://yourhost.publicaddress.com/cancel|_Cancels Incident_|
|`/hellper_update_dates`|https://yourhost.publicaddress.com/dates|_Updates the dates for an incident_|

The first command `/hellper_incident` can be use at any channel and/or conversation on Slack. It will open a pop-up for the user to set and start an Incident, creating the channel, meeting room link and post-mortem doc.

The remaining commands must be used only on the Incident's channel since they act on the specific incident that is open.


### Metrics
This metrics came from `metrics` view table, they are calculated by the following formulas:

| Metric | Description | Formula |
| --- | --- | --- |
| **start_ts** | Date and time when the incident is started | Date and time in UTC from db |
| **identification_ts** | Date and time when the incident is identified | Date and time in UTC from db |
| **end_ts** | Date and time when the incident is resolved | Date and time in UTC from db |
| **acknowledgetime** | Time To Acknowledge | `identification_ts` - `start_ts` |
| **solutiontime** | Time To Solution | `end_ts` - `identification_ts` |
| **downtime** | Time in an incident | `end_ts` - `start_ts` |
| **MTTA** | Mean Time To Acknowledge | `total acknowledgetime` / `total incidents` |
| **MTTS** | Mean Time To Solution | `total solutiontime` / `total incidents` |
| **MTTR** | Mean Time To Recovery | `total downtime` / `total incidents` |

## Contributing
Thanks for being interested in contributing! We’re so glad you want to help! Please take a little bit of your time and look at our [contributing guidelines](/docs/CONTRIBUTING.md). All type of contributions are welcome, such as bug fixes, issues or feature requests.

## Code of Conduct
Everyone interacting in the Hellper project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](/docs/CODE_OF_CONDUCT.md).

## Need help?
If you need help with Hellper, feel free to open an issue with a description of the problem you're facing.

## License
The Hellper is available as open source under the terms of the [MIT License](LICENSE).
