import os
import requests
from requests.auth import HTTPBasicAuth
import json
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

        print(resp)

        return resp.json
