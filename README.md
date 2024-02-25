# Go Muse project for CLI App

## commands available in muse

* config
* make
  * js *// not implement yet*
  * ts *// not implement yet*
  * py *// not implement yet*
  * java *// not implement yet*
  * go *// not implement yet*
* version

## TODOS

### Root and Version

- [x] Add root command with subcomands for version, config and make

- [ ] Add funcionality for root subcommand called version that displays the version number if its posible that call an api to get version number and displays if is updatable

### Config

- [x] Add functionality for root subcomand called config which adds new items to db into json file and validates that the args have been two
- [x] Create JSON file with go for save the aliases and paths
- [x] Validate path arg in config command (may or may not exist)
- [x] Validate alias arg in config command between 2 to 10 chars and it musb be unique
- [x] Insert into aliases JSON file with go that creates Item for config command

### Make

- [x] The alias flag must be more relevant than the output flag
- [x] Validate that the alias flag exist. If only the flag output was given, it does not matter that the route does not exist
- [x] Validate the flag name only must contains letters and underscores
- [x] Create subcomands for javascript, typescript, java, python, and go.
- [x] Java command will only be available using spring boot. Create the structure according to the data provided
- [ ] JavaScript and TypeScript will only be available for frontend with vite.
- [ ] Go in progress - !NOT DEFINED¡
- [ ] Python command will only be available with venv enviroment. And implemented
- [ ] Create a directory based on inputs

### List 

- [x] Make command for list
- [x] Create flag template to list available templates


*CATCH ERRORS ON ALL COMANDS*