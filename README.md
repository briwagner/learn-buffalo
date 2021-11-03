# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

## Part 4: using controllers with our models

* CLI steps to create an action and handler function
* modify the handler to return some data
* create a grift task to seed data in the DB
* testing controllers with ActionSuite, using test fixtures to setup test data
* pop Eager() method to load model associations

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

### Going Further

* add index listing page for all blogs

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
