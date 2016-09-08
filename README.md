# scraping-targets
An example website used as a target for web scraping during the [Node.js Web Scraping](https://github.com/paulzerkel/pres-node-web-scraping) presentation.

## Overview
This is an extremely simple web application built using Go. The app is not intended to be run on a public facing server so buyer beware. Compile and start the application with `go build && ./scraping-targets`. The application will start up at [`http://localhost:8080`](http://localhost:8080) and output basic logging to the console.

## Misc
The site makes use of Bootstrap and Knockout. The `/product-data` endpoint was intentionally build with a long pause in order to demonstrate benefits of web scraping with a full web client.