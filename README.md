# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

> Updated Dec. 2023 to work with `go:embed` directive and other changes for Buffalo version <a href="https://github.com/gobuffalo/cli/releases/tag/v0.18.14">18.14</a>

## Part 2: create a user model, add fields and methods, and test it

* Models as part of the MVC paradigm
* Buffalo patterns inspired by Ruby on Rails and the influence of test-driven development
* configure database connections in database.yml file
* use the Buffalo cli to generate a model; specify additional fields and field types
* running buffalo test command
* writing tests using the ModelSuite
* testing with the database

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
