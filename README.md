# discord-notifier

Go based lambda notfier to send messages to discord.

## Prerequists

go 1.x
serverless

## Build

`make`

## Deploy

DISCORD_WEBHOOK_URL must be set to a valid discord webhook

`sls deploy`

## Test

`sls invoke -f discord-notifier`

This you send a message to your channel
