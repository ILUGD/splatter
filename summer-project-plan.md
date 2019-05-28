# ILUGD Summer Project Plan

What do we need?
 1. A server to serve the tool as an API
 2. A frontend to draw the templates
 3. A format to store the templates with assets and the config
 
### Deliverable 1
 - Tests are missing, so completing the tests
 - The positions are hardcoded to an extent
 - Probably also add CI to build and test the stuff
 - Have a dockerfile, it would help when we will the serving it as an API
 
### Deliverable 2
 - Use `net/http` to serve the existing application
 - API for sending a config file and as a response get an image
 - A simple frontend for uploading the config and obtaining the result in the browser

### Deliverable 3
 - More work on the frontend, creating a canvas
 - With the canvas we can interactively generate a config

### Deliverable 4
 - Generate `.splat` files to contain both the config and the assets
 - Either as an archive of `SVG` and `json` but parsing SVG can sometimes be tricky with the possibilities it offers 
 - Or another option is to generate a binary file using `encoding/gob` 
 
Languages : Golang, JavaScript
External Libraries : Cairo for Golang, no idea about the JS part though
