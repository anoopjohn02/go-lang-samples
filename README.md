# go-lang-samples

A goLang sample project that represent following features.

- APIs using Gin
- Security with Middlewares
- Mongo DB connectivity
- MQTT Connectivity
- Caching
- Tests with mock

# Features

This microservice is implementation of simple fleet/edge device in IoT industry.

- Measurement: API to receive device measurement and save to mongo db. The device should be authenticated.
- Alert: API to send alerts to the MQTT and saved to the local mongo DB. The device should be authenticated.
- Cache the device profile for a while.
