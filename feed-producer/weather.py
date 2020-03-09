from openweather_api import OpenWeatherAPI
from repository import Repository


# implementa il tag "weather" con persistenza su repository -> redis
class Weather:

    def __init__(self, apikey, cities):
        self.cities = cities
        self.api = OpenWeatherAPI(apikey)
        self.repo = Repository()

    def forecast(self):
        for city in self.cities:
            self.repo.push("weather|" + city, self.api.three_hour_forecast(city))
