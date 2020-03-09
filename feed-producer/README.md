## Feed Producer

## Test JSON weather scraping and storing in Redis (development environment)
1. Open the __deploy directory
2. Set your configuration parameters in the docker-compose-dev.yml file (PRODUCER_WEATHER_CITIES, PRODUCER_WEATHER_TIME)
3. From __deploy directory, open a terminal and launch the application (details in __deploy [README.md](https://github.com/AlessandroSpallina/SmartFeed/blob/master/__deploy/README.md))
4. Open a second terminal and run:
```
docker exec -it redis  redis-cli
```
5. While in redis-cli, run:
```
monitor
```
6. You will see your list of dictionaries being stored periodically (period decided by the PRODUCER_WEATHER_TIME parameter setted before)

## Demo

