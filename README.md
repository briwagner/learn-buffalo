# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

## Part 3: models and associations, one-to-one, one-to-many, many-to-many

* one-to-one relation: User -> Address
* one-to-many relation: User -> Blogs
* using join tables to manage many-to-many relations
* many-to-many relation: Blogs <-> Tags
* manage child entities with pop's Eager() method, or using model methods

### Installation

Modify the database.yml file in the root folder to use the database type, database name and password to match your machine.
The current file assumes:
* type: mariadb
* db name: modelsdemo for dev, and modelsdemo_test for test
* db user: buffalo
* password: 'password'

### Going Further

* foreign-key constraint violation; allow child entities to have NULL for parent

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
