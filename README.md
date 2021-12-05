# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

## Part 7: Authenication Plugin, Middleware

* user authentication patterns
* Plugins in Buffalo to generate boilerplate code
* add email, password_hash fields to User, to use with authentication
* using Middleware to check for login status, and fetch data for use in "next" handler

## Next Steps

* change model tests to reflect new fields on User
* move DB setup steps from model tests into fixtures
* add tests for new login/logout routes

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
