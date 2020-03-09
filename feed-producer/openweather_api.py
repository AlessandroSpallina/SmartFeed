import requests
import json
from utils import extract_values

# @finfme : USA QUESTO FORMATO DI CLASSE PER LE VARIE API
class OpenWeatherAPI:

    def __init__(self, apikey):
        self.apikey = apikey
        self.url = 'https://api.openweathermap.org/data/2.5'

    def __url(self, path):
        return self.url + path

    def three_hour_forecast(self, city):
        params = {
            'q': city,
            'mode': 'json',
            'appid': self.apikey
            }

        resp = requests.get(self.__url('/forecast'), params=params)

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
