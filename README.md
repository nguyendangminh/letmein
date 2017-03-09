# Slack, let me in!

Send Slack invitation automatically.

![Slack, let me in](https://github.com/nguyendangminh/letmein/blob/master/screenshot.png)

## Install
```
$ mkdir slack
$ cd slack
$ export GOPATH=$(pwd)
$ go get github.com/nguyendangminh/letmein
```

## Run
$ SLACK_TOKEN=your_slack_token PORT=your_port TEAM_NAME=your_team_name $GOPATH/bin/letmein
