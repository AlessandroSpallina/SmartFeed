## Identity Node
[![Go Report Card](https://goreportcard.com/badge/github.com/AlessandroSpallina/SmartFeed)](https://goreportcard.com/report/github.com/AlessandroSpallina/SmartFeed)

Read our beautiful [API Docs](https://documenter.getpostman.com/view/13959997/TWDXpHP6)

## Architecture
<p align="center">
  <img src="https://spee.ch/d/identity-node.png">
</p>

## Test REST stuff with postman (development environment)
1. Install [Postman](https://www.postman.com/)
2. Import the postman method collection (click on run in postman) [HERE](https://documenter.getpostman.com/view/13959997/TWDXpHP6)
3. Choose methods to test by GUI.

## Test MQTT stuff with MQTT Explorer (development environment)
1. Install [MQTT Explorer](http://mqtt-explorer.com/)
2. Connect to the broker (localhost:1883)
3. Publish a request in the welcome topic (identity/welcome) following the format:
```
{
  "username": <previouslyRegisteredUser>,
  "response_topic": <aTopicName>
}
```
4. You will get a reply in the "response-topic" requested following the format:
```
[
  {
    "tag": "weather",
    "args": {"city": ["catania", "giarre"]}
  }
]
```

## Demo
[![Alt text](https://spee.ch/3/identity-node-demo2.jpg)](https://spee.ch/d/identity-node-demo-v0.webm)
