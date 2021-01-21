# Doggo 🐶

*This tool is tied to this project's context (for now)*

Doggo is an helper CLI tool for the Owly project.

## Install

When you are in this directory, just type `go install doggo` and you should be able to use `doggo` from your command line.

## Functionnalities

Type `doggo --help` to see all available commands and functionalities


### `connect`

**Desc:** Opens an Evans session at desired port, using a user's token from provided username and password


So, all you need, is the username and the password of the user you want the token from.

Then :

`doggo connect -U <username> -P <password> -p <port>`


### `docker_clean`

**Desc:** Clean volume, container, system cache out of Docker.

Basically does this : `docker system prune -f && docker container prune -f && docker volume prune -f`

## Additional notes

Of course, you need : 
- evans 
- docker

to be already installed and set in your PATH :)



*There u (dog)go ! 🐶*

by [@AppliNH](https://github.com/AppliNH)