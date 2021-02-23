## Feed Provider

## Architecture
<p align="center">
  <img src="">
</p>

## Test REST stuff with postman (development environment)
1. Install [Postman](https://www.postman.com/)
2. Import the postman method collection (click on run in postman) [HERE]()
3. Choose methods to test by GUI.

## Test MQTT stuff with MQTT Explorer (development environment)
1. Install [MQTT Explorer](http://mqtt-explorer.com/)
2. Connect to the broker (localhost:1883)
3. Publish a request in the welcome topic (provider/welcome) following the format:
```
{
  "query": [
                {
                    "tag" : <tagName>,
                    "args" : ['value1', 'value2']
                }
           ],
  "response_topic": <aTopicName>
}
```
4. You will get a reply in the "response-topic" requested following the format:
```
[
  {
    "query": {
      "tag": "weather",
      "args": [
        "catania"
      ]
    },
    "data": [
      {
        "temp": 286.73,
        "feels_like": 285.89,
        "temp_min": 286.17,
        "temp_max": 286.73,
        "pressure": 1033,
        "humidity": 75,
        "description": "clear sky",
        "dt_txt": "2021-02-23 18:00:00"
      }, ...
    ]
  },
  {
    "query": {
      "tag": "weather",
      "args": [
        "misterbianco"
      ]
    },
    "data": [
      {
        "temp": 285.76,
        "feels_like": 284.94,
        "temp_min": 284.26,
        "temp_max": 285.76,
        "pressure": 1033,
        "humidity": 76,
        "description": "clear sky",
        "dt_txt": "2021-02-23 18:00:00"
      }, ...
    ]
  }
]
```