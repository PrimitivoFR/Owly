# Owly
🦉  An open-source chat app project with a self-hostable architecture


## Run it

Take care to run the **test mode** and check that everything runs fine with your new developments before commiting.

### Dev env

Run everything:
`docker-compose -f docker-compose.yml -f docker-compose.devandtest.yml  -f docker-compose.dev.yml up`

Restart a specific service:
`docker-compose -f docker-compose.yml -f docker-compose.dev.yml restart SERVICE_NAME`

#### Dialing with EVANS CLI

You first have to pass the token before making any request !

`header set authorization="BLABLA_TOKEN"`

#### About the client in dev mod

The owly-client is not included anymore in the dev env stack.
It was resulting in making the whole dev env highly slower, and was impacting negativly the workflow.

In brief, it was way to heavy to load.
For now, you'll have to run apart:

`npm i && npm start` to start the client.



### Test mode

Run everything:
`docker-compose -f docker-compose.yml -f docker-compose.devandtest.yml -f docker-compose.test.yml up`

### Local Prod mode

Allows to simulate in a local something approaching a real production stage.

Please, execute `server/dockerBuildAll.sh` first, it's **needed**.

`docker-compose -f docker-compose.yml -f docker-compose.prod-local.yml up`.

### Prod mode

Prod mode is not quite ready yet. We'll pull images from the Github Container Registry.