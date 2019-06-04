### Getting started
 - Install Go, instructions for that can be found [here](https://golang.org/doc/install)
 - Install cairo development libraries, instructions for that can be found [here](https://www.cairographics.org/download/)
 - Do a `go get -u -v github.com/ILUGD/splatter/...`
 - Fork the repo on github
 - Go to `$GOHOME/src/github.com/ILUGD/splatter/`
 - Add the url of your fork, `git add remote <your username> <your fork's url>`
 - Switch to a new branch
 - Commit the changes
 - Push to your fork, `git push <your username> <your branch name>`
 - Create a Pull Request

### Styles and naming conventions
 - run the code through `goimports` before committing
 - structs/functions should be documented in the following manner
```
//Function/Struct Name <2 spaces> What it does
Function/Struct definition
```
 - Try to keep the lines short typically to 75 - 80 chars at most
 - Naming should only be done in camelCase

### Opening an Issue
 - Issue should start with the package associated with it., followed by a colon and short description*
**Example** : `parser: Variable a is not being parsed`
 - Append [Issue] if it is a problem, [Request] if it is a feature request to the description
 - Don't use punctuation marks at the end of the title

### Submitting Patches
 - Comment on an issue before starting working on it
 - If it is a feature make a issue on it
 - Start working after you are assigned to the issue
 - Commit Message should begin with the package you have worked on followed by a colon and short description, just like Issue titles *
 - Always append the description with **Fix:** #issue_number
 - Fork and Commit your changes in a separate branch

If you are still not sure about the guidelines, you can ask us on our IRC/matrix channel, Telegram group or Mailing Lists.

