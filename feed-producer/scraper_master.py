#componente che si occupa di gestire lo scraping degli interessi tramite richieste API REST
#l'obiettivo è fare in modo che redis riceva periodicamente dati 
#parametri di configurazione per gli scraper sono definiti nel file conf.yml
#

import yaml


with open("conf.yml", 'r') as ymlfile:
    cfg = yaml.load(ymlfile)
