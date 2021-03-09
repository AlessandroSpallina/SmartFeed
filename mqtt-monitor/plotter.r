makePlot <- function (png_save_path= "plot.png") {
  a <- read.csv(file=file.path("mqtt-monitor", "mqtt_messages.csv"), sep=";")
  # voglio un plot multiclasse (una per topic), il tempo nelle x e i byte nelle y
  a$timestamp <- as.POSIXct(a$timestamp, origin="1970-01-01")
  # voglio definire i topic come classi (colore diverso nel plot), serve renderli factor
  a$topic <- as.factor(a$topic)

  png(file.path("mqtt-monitor", png_save_path))
  plot(x=a$timestamp, y=a$byte_received, col=a$topic, xlab="time", ylab="byte received", pch=1:nlevels(a$topic))
  legend("topright", pch=1:nlevels(a$topic), title="topics", legend=unique(a$topic), col=a$topic)
  dev.off()
}

makePlot()