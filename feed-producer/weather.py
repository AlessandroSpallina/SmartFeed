from openweather_api import OpenWeatherAPI
from provider_api import ProviderApi


# implementa il tag "weather" con persistenza su microservizio feed-provider
class Weather:

    def __init__(self, apikey, cities):
        self.cities = cities
        self.api = OpenWeatherAPI(apikey)
        self.provider = ProviderApi()

    def forecast(self):
        for city in self.cities:
            self.provider.produce('weather', [city], self.api.three_hour_forecast(city))
            # self.repo.push("weather|" + city, self.api.three_hour_forecast(city))
