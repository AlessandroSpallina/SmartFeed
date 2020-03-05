#componente che fa scraping delle previsioni meteo dal sito openweather.org
#
#
#

import yaml
import requests
import os
import json
import redis

#recupero i valori di configurazione dal file conf.yml e le variabili d'ambiente dal docker-compose.yml
#with open("conf.yml", 'r') as ymlfile:
   # cfg = yaml.load(ymlfile)


#API_KEY='79c9118b24b5b9dc61998d484842f0dc'

#importo le variabili d'ambiente
API_KEY=os.environ.get('API_KEY', 'aaaaaaaaaaaaa')
city=os.environ.get('city', 'londra')

#costruisco il payload da passare alla request per ottenere il json
payload={'q':city, 'mode':'json', 'appid':API_KEY}

response = requests.get('https://api.openweathermap.org/data/2.5/forecast', params=payload)

#verifica che l'url sia corretto
print(response.url)

#converto i json ricevuto in un dizionario
json_data=json.loads(response.text)


