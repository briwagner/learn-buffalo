# Learn Buffalo

This repo contains the code used in the Learn Buffalo video course.

> Updated Dec. 2023 to work with `go:embed` directive and other changes for Buffalo version <a href="https://github.com/gobuffalo/cli/releases/tag/v0.18.14">18.14</a>

Sections Completed:

1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-1">Part 1</a>: new project, static page, and route parameters
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-2">Part 2</a>: create a user model, add fields and methods, and test it
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-3">Part 3</a>: models and associations, one-to-one, one-to-many, many-to-many
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-4">Part 4</a>: using controllers with our models
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-5">Part 5</a>: HTML templates, partials and helpers
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-6">Part 6</a>: Forms, form_for() helper, form() helper, manually parse and grab form values from the request
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-7">Part 7</a>: Middleware and user authentication with Buffalo Plugins
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-8">Part 8</a>: Third-party Integrations and Events
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-9">Part 9</a>: CLI operations to build a production build, and other tasks
1. <a href="https://github.com/briwagner/learn-buffalo/tree/part-10">Part 10</a>: Develop and Deploy with Docker, using a cloud hosted database, or with docker-compose

## Installation

Fork the project (if you want to commit and push your own changes).
Each part listed above has a matching branch. The branch for part-6, for example, is the end-state of the project code at the completion of video part-6. If you want to follow along with the part-6 video, you can `checkout part-5` and start making changes.
After cloning the repo to your local machine, `git checkout {branch-name}` to view the part.

When changing branches, use these handy buffalo commands to get a clean project state:
* `buffalo pop reset` - this will wipe the database and rebuild with the migrations
* `buffalo task db:seed` - part-4 and beyond include a DB seed task; this will load the relevant data into your dev database
* `buffalo test` - run the tests to make sure everything is working before you get started
* `buffalo dev` - run the dev server to see the site in your browser

## Buffalo Project

The project site [http://gobuffalo.io](http://gobuffalo.io) has lots of great documentation there.
