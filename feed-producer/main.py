import schedule
import os
import pyfiglet
import time
from openweather_api import OpenWeatherAPI
from repository import Repository

def get_config_from_env():
    return {
        "OPENWEATHER_APIKEY": os.environ['PRODUCER_OPENWEATHER_APIKEY'],
        "WEATHER_CITIES": os.environ['PRODUCER_WEATHER_CITIES'],
        "WEATHER_STATES": os.environ['PRODUCER_WEATHER_STATES'],
        "WEATHER_TIME": os.environ['PRODUCER_WEATHER_TIME']
    }

def main():
    config = get_config_from_env()

    print(pyfiglet.figlet_format("- RFP -")) # RFP codename for REST Feed Producer

    tic_counter = 0
    weather = OpenWeatherAPI(config['OPENWEATHER_APIKEY'], config['WEATHER_CITIES'].lower(), config['WEATHER_STATES'])

    #scheduling delle api (v0.1)
    while True:
        time.sleep(1)
        if tic_counter % int(config['WEATHER_TIME']) == 0:
            previsions=weather.forecast_hourly()
            repo = Repository(config['WEATHER_CITIES'])
            repo.push_repo('weather', previsions)


main()
