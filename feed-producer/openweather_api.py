import os
import requests
from requests.auth import HTTPBasicAuth
import json
import datetime
#import csv

# @finfme : USA QUESTO FORMATO DI CLASSE PER LE VARIE API
class OpenWeatherAPI:

    def __init__(self, apikey, cities, states):
        self.apikey = apikey
        self.cities = cities
        self.states = states
        self.url = 'https://api.openweathermap.org/data/2.5'

    def _url(self, path):
        return self.url + path

    def forecast_hourly(self):
        params = {
            'q': self.cities,
            'mode': 'json',
            'appid': self.apikey
            }

        weather = {
            "temp": 'temp',
            "feels_like": 'feels_like',
            "temp_min": 'temp_min',
            "temp_max": 'temp_max',
            "pressure": 'pressure',
            "humidity": 'humidity',
            "description": 'description',
            "dt_txt": 'dt_txt'
            }
            
        resp = requests.get(self._url('/forecast'), params=params)

        weather_dict = resp.json()["list"]
        weather_list=[]
        main_list = ['temp', 'feels_like', 'temp_min', 'temp_min', 'temp_max', 'pressure', 'humidity' ]

        for items in weather_dict[:3]:
            weather['dt_txt'] = items['dt_txt']
            for key, value in items['main'].items():    #per ogni coppia nel dizionario items['main']
                if key in main_list:
                    weather[key] = value
            for key, value in items['weather'][0].items(): #per ogni coppia nel dizionario items['weather']
                if key == 'description':
                    weather['description'] = value
            #for key, value in items['dt_txt'].items():
            weather_list.append(weather)
        return(weather_list)
