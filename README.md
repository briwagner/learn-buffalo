# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

## Part 9: Build for Production and Deploy

* `buffalo build` command and flags (-o, -v, --clean-assets)
* production mode versus development mode
* enabling Buffalo TLS server
* checklist for deploying on a remote server (i.e. Ubuntu linux)
* systemd configuration
* nginx configuration
* securing sensitive routes

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
