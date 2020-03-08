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
    #print(config['OPENWEATHER_APIKEY'])

        #schedule.every(confi.WEATHER_TIME).seconds()

    tic_counter = 0
    weather = OpenWeatherAPI(config['OPENWEATHER_APIKEY'], config['WEATHER_CITIES'], config['WEATHER_STATES'])
    #print(weather.forecast_hourly())
        #weather.forecast_hourly()
        #schedule.every(config['WEATHER_TIME']).seconds.do(weather.forecast_hourly)

    #schedule.every(10).seconds.do(weather.forecast_hourly)
    #schedule.every(1).minutes.do(test)

    #print(pyfiglet.figlet_format("-------"))

    #scheduling delle api (v0.1)
    #while True:
	 #   time.sleep(1)
	  #  if tic_counter % WEATHER_TIME == 0:
    previsions=weather.forecast_hourly()
    repo = Repository(config['WEATHER_CITIES'])
    repo.push_repo('weather', previsions)


main()
