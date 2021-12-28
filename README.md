# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

## Part 8: Third-party Integrations and Events

* storing credentials in .env file
* connecting to MongoDB, Slack webhooks, Sendgrid API
* Buffalo events and event-listeners
* adding a health check in the app launch process
* emitting events with payload data
* adding event listeners in the app.go file and custom packages

## Next Steps

* testing connections to third-party services
* extending Mongo to read other collections; storing data

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
