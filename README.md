# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

> Updated Dec. 2023 to work with `go:embed` directive and other changes for Buffalo version <a href="https://github.com/gobuffalo/cli/releases/tag/v0.18.14">18.14</a>

## Part 6: HTML forms amd helpers

* form_for() used with a model or struct
* form() helper to manually build form elements
* generate a select tag from a map of values
* add SelectValue and SelectLabel to satisfy selectable interface
* build handler to accept form data
* Bind() method to decode form into struct
* manually process the form to process individual fields

## Next Steps

* add tests for the BlogsNew and BlogsCreate routes
* modify the Tags field to use checkboxes

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
