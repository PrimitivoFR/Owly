# Owly
🦉  An open-source chat app project with a self-hostable architecture


## Run it

Take care to run the **test mode** and check that everything runs fine with your new developments before commiting.

### Dev mode

You'll need `minikube` and `kubectl`

Run everything:
- `make dev`
- open another terminal tab, and type `minikube ip`. Match this ip with domain `owly.dev` in your /etc/hosts (eg: `192.168.49.2         owly.dev`)
- open another terminal tab and run `skaffold dev`
- the head up to https://owly.dev
- **NOTE:** if your browser doesn't let u pass, just type `thisisunsafe` on your keyboard.

Port-forwarding for dev: 
- `make port-fwd-all`


#### Dialing with EVANS CLI

Some service require a token.
First, create yourself an account using evans-cli :
- `evans -r repl -p 50054`

Then, memorize your username and password, so you can use `doggo` to make request to other services, using your account token.
For more informations, go and see the doggo's Readme at `.devtools/doggo`


### Test mode

Run tests:
- `make test`

### Prod mode

Prod mode is not quite ready yet. We'll pull images from the Github Container Registry.
