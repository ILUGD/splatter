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
 - This would use `VueJS` for the web app
 - It would generate `.splat` files
 - A custom API would accept the splat files and interface with the CLI to generate an image
 - The frontend would allow the user to download the image

### Deliverable 4
 - Generate `.splat` files to contain both the config and the assets
 - Either as an archive of `SVG` and `json` but parsing SVG can sometimes be tricky with the possibilities it offers 
 - Or another option is to generate a binary file using `encoding/gob` 
 
**Languages** : Golang, JavaScript

**External Libraries** : Cairo for Golang, no idea about the JS part though

##### Why Golang?
 - This is one of the language I am most comfortable with, so gate-keeping for me(@hellozee) will be easy
 - A part of the project is already done in it, so less work
 - The whole toolchain is bundled into binary and lesser permutation of approaches
 - The standard library is most often more than enough

##### Why Cairo?
 - One of most popular **drawing** libaries out there
 - Device independent, can directly output to SVG/PDF, along with several image formats
 - Bindings exists for most of the languages
 - The only alternative looks like [skia](skia.org), which is probably limited to C++ and lacks community support like cairo
