# Slack, let me in!

Send Slack invitation automatically.

![Slack, let me in](https://github.com/nguyendangminh/letmein/blob/master/Screen%20Shot%202016-05-18%20at%203.43.45%20PM.png)

## Install
```
$ mkdir slack
$ cd slack
$ export GOPATH=$(pwd)
$ go get github.com/nguyendangminh/letmein
```

## Run
$ SLACK_TOKEN=your_slack_token PORT=your_port TEAM_NAME=your_team_name $GOPATH/bin/letmein
