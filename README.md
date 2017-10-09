# Go + Docker + DigitalOcean Example

A minimal example demonstrating how to create a dockerized Go app that can easily be deployed as a DigitalOcean droplet.

## Structure

```
/src/...
/docker-compose.yml
/Dockerfile
```

### App
The app source is placed under `/src`

### Dockerfile
- The dockerfile is based off of the official `golang` based image, which contains the desired go environment.
- Different variants are example. For example, for a specific version of go, use `golang:1.6` instead of `golang:latest`.
- For more information see the [golang](https://hub.docker.com/_/golang/) docker image.

### docker-compose
The docker-compose configuration sets up the `app` service and a simple `nginx-proxy` service for routing.

The `nginx-proxy` service automatically links up any other services that have a `VIRTUAL_HOST` environment variable, and forwards any requests that match the routes.

Both services use `restart: always` to automatically restart in case the machine restarts.

Log history is capped at `100m` to prevent running out of disk space 😅 (I learned that the hard way once).

## Testing
To run the services locally
```
docker-compose build && docker-compose up
```

To run the services locally (detached)
```
docker-compose build && docker-compose up -d
```

To see the status:
```
docker-compose ps
```

To test the app:
```
curl -H 'Host: foobar.example.com' localhost
```

 Note that we have to supply an explicit host header since we're testing locally. We could also add another `VIRTUAL_HOST` to the service for `localhost` to receive requests from `http://localhost/`.

## Deploying
Deploying to DigitalOcean is pretty straightforward.

### 1) Create a new droplet
- `Create` > `Droplet`
- `One-click apps`
- `Docker 17.05.0-ce on 16.04` (or newer)
- Choose desired `size`
- Choose desired `region`
- Select `Monitoring`
- Choose desired `hostname`

### 2a) Deploy using User Data
When creating the droplet, enable `User Data` and enter your script:
```
#!/bin/bash
git clone https://<token>@<repo> ~/app
cd ~/app && docker-compose build && docker-compose up -d
```

Note: Logs for `User Data` cloud init scripts can be found here: `/var/log/cloud-init-output.log`.

### 2b) (Re-)Deploy using SSH
1. SSH into the droplet
  - `ssh root@<ip>`

2. Clone the source repo:
  - `git clone <repo>`

3. Deploy
  - `cd app && docker-compose build && docker-compose up -d`

4. Optionally, create monitoring alerts in DigitalOcean

5. To re-deploy, repeat steps 2-4
