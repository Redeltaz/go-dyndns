# go-dyndns

Dynamic DNS writtent in Go, only work with [Scaleway](https://www.scaleway.com) for the moment

## How to use

This DynDNS work as a simple executable that when executed, update you subdomain IP to you new public IP, you can use it how you want (loop in other script, scheduled lamdba...) but for me I prefer to use it with a cronjob on my server.

So you can either only build the executable or run it in a premade docker container with pre-confiruged cronjob. Here is how you can do it :

First of all you need to have Docker latest version installed on your machine, or at least a version that support every line of the [docker-compose.yml](./docker-compose.yml) file

Then copy paste the [.env.example](./.env.example), name it `.env` and fill the env variable such as :

- SW_SECRET_KEY: *The secret key you get when generating an API Key on Scaleway*
- DOMAIN_NAME: *The domain name used*
- SUBDOMAIN_NAME: *The subdomain linked to your public IP address*
- SCHEDULER_CRON: *The cron expression for the scheduled excecution, this env variable is not needed if you only want to build the executable without running it in the preconfigured image I made*

```bash
mv .env.example .env
```

### Build

To only build the executable you can execute this command :
```bash
docker compose up --build -d builder
```

The executable will be available in the `./bin` folder

### Run

To run the dynamic DNS in a preconfigured container with cronjob configured you can execute this command :
```bash
docker compose up --build -d runner
```

The container will run as itself and execute the command following your scheduled cron expression (or every 5 minutes by default if you didn't add an expression). The logs will be available on your machine in : `/var/log/go-dyndns/go-dyndns.log`
