import os
import requests
from requests.auth import HTTPBasicAuth
import json
from parse_api import extract_values
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
            
        resp = requests.get(self._url('/forecast'), params=params)

        weather_list= []
        #weather_list = extract_values(resp, key) 
        #mi torna una lista con tutti i valori di key nell'oggetto resp
        #quindi se metto 'temp' avro' una lista con tutte le temperature misurate (ne voglio solo 3)

        key_list = ['temp', 'feels_like', 'temp_min', 'temp_min', 'temp_max', 'pressure', 'humidity', 'description', 'dt_txt' ]
        temp_list = []
        x = 0

        for x in range(0, 3):
            weather = {}
            for key in key_list:
                temp_list = extract_values(resp.json(), key)
                weather[key] = temp_list[x]
            weather_list.append(weather)

        return weather_list

