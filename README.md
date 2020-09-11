# Owly
🦉  An open-source chat app project with a self-hostable architecture


## Run it

### Dev env

Run everything:
`docker-compose -f docker-compose.yml -f docker-compose.dev.yml up`

Restart a specific service:
`docker-compose -f docker-compose.yml -f docker-compose.dev.yml restart SERVICE_NAME`

#### About the client in dev mod

The owly-client is not included anymore in the dev env stack.
It was also resulting in making the whole dev env highly slower, and was impacting negativly the performances of the dev env.

In brief, it was way to heavy to load, and for now, we can't figure it out how to fix it.
In the meantime, you'll have to run apart:

`npm i && npm start` to start the client.

Take care to run the **test mode** and check that everything runs fine with your new developments before commiting.


### Test mode

Run everything:
`docker-compose -f docker-compose.yml -f docker-compose.dev.yml -f docker-compose.test.yml up`