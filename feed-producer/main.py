import schedule
import os
import pyfiglet
import time
#from openweather_api import OpenWeatherAPI
#from repository import Repository
from weather import Weather

def get_config_from_env():
    return {
        "WEATHER_APIKEY": os.environ['PRODUCER_WEATHER_APIKEY'],
        "WEATHER_CITIES": os.environ['PRODUCER_WEATHER_CITIES'].lower().split("|"),
        #"WEATHER_STATES": os.environ['PRODUCER_WEATHER_STATES'],
        "WEATHER_TIME": int(os.environ['PRODUCER_WEATHER_TIME'])
    }

def main():
    config = get_config_from_env()

    print(pyfiglet.figlet_format("- RFP -")) # RFP codename for REST Feed Producer

    tic_counter = 0
    weather = Weather(config['WEATHER_APIKEY'], config['WEATHER_CITIES'])

    #scheduling delle api (v0.1)
    while True:
        time.sleep(1)
        if tic_counter % config['WEATHER_TIME'] == 0:
            weather.forecast()
        tic_counter += 1


if __name__ == "__main__":
    main()
