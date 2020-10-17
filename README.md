# emoji-race
Display in real time the most popular K emojis from twitter.
The application is inspired by https://observablehq.com/@mbostock/twitter-emoji-race

## API service

The [API service](https://github.com/efekarakus/emoji-race/tree/mainline/api) is a load balanced web service to display results.

## Tracker service

The [Tracker service](https://github.com/efekarakus/emoji-race/tree/mainline/tracker) is a backend service that is subscribed to a stream of Twitter emojis 
from [Emojitracker](https://github.com/mroth/emojitracker).

## Deploy using AWS Copilot
This application is deployed using the [AWS Copilot CLI](https://github.com/aws/copilot-cli).
The resulting artifacts are under the [copilot/](https://github.com/efekarakus/emoji-race/tree/mainline/copilot) directory
