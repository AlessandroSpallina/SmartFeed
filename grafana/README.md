## Grafana Dashboards

1. Go to grafana url, default http://localhost:80
2. Login on grafana, default credentials are admin admin
3. Add data source prometheus, default url http://prometheus:9090
4. [Import](https://grafana.com/docs/grafana/latest/reference/export_import/#importing-a-dashboard) the dashboard config file present in this directory
5. Don't forget to select the right data source (prometheus!)

Enjoy some statistics about the MQTT broker!

![](https://spee.ch/7/mqtt-dashboard.jpg)
