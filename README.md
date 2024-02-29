# Go Muse project for CLI App

## commands available in muse

* config
* make
  * js *only vite*
  * ts *only vite*
  * py *// not implement yet*
  * java
  * go *// not implement yet*
* version

## TODO'S

### Root and Version

- [x] Add root command with subcommands for version, config and make

- [ ] Add functionality for root subcommand called version that displays the version number if its possible that call an api to get version number and displays if is updatable

### Config

- [x] Add functionality for root subcommand called config which adds new items to db into json file and validates that the args have been two
- [x] Create JSON file with go for save the aliases and paths
- [x] Validate path arg in config command (may or may not exist)
- [x] Validate alias arg in config command between 2 to 10 chars and it must be unique
- [x] Insert into aliases JSON file with go that creates Item for config command

### Make

- [x] The alias flag must be more relevant than the output flag
- [x] Validate that the alias flag exist. If only the flag output was given, it does not matter that the route does not exist
- [x] Validate the flag name only must contains letters and underscores
- [x] Create subcommands for javascript, typescript, java, python, and go.
- [x] Java command will only be available using spring boot. Create the structure according to the data provided
- [x] JavaScript and TypeScript will only be available for frontend with vite.
- [ ] Go create projects with framework like Fibber, Gin, etc.
- [ ] Python command will only be available with *venv environment*. And implemented
- [x] Create a directory based on inputs
- [x] When create directory, warn the user about the existence of the directory

### List 

- [x] Make command for list
- [x] Create flag template to list available templates
- [ ] Add list of commons dependencies for java spring boot

### Persistent

- [ ] save the project data for simplify the creation of the subsequent projects

*CATCH ERRORS ON ALL COMMANDS*

*bug in utilities: multiple selection, single selection. when changing the cursor with right arrows, the cursor is not located at the start of the element*

*the bug is for the size of screen*
