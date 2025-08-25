<div align="center">
    <img src="docs/logo/logo.jpg" width="500px" align="center" alt="sf metadata linter logo" />
</div>
</br>
<div align="center">
    SkyStats is an application to retrieve, store, and display interesting aircraft ADS-B data received via an SDR.
</div>

## Overview

* [Go](https://go.dev/) app with [PostgreSQL](https://www.postgresql.org/) database and [Svelte](https://svelte.dev/) + [DaisyUI](https://daisyui.com/) front end
* ADS-B data is received via [adsb-ultrafeeder](https://github.com/sdr-enthusiasts/docker-adsb-ultrafeeder) / [readsb](https://github.com/wiedehopf/readsb), running on a Raspberry Pi 4 attached to an SDR + aerial ([see it here!](docs/setup/aerial.jpg))
* The application consumes aircraft data from the readsb [aircraft.json](https://github.com/wiedehopf/readsb-githist/blob/dev/README-json.md) file
* A [gin](https://gin-gonic.com/) API surfaces information from the postgres database to the web frontend
* Registration & routing data is retrieved from the [adsb-db](https://github.com/mrjackwills/adsbdb) API
* "Interesting" aircraft are identified via a local copy of the [plane-alert-db](https://github.com/sdr-enthusiasts/plane-alert-db)

## Features
* "Above Me" - live view of 5 nearest aircraft with routing information
* Total aircraft seen (past hour, day, all time)
* Total aircraft with route data
* Unique Countries
* Unique Airports
* Top Airlines
* Top Airports (Domestic, International)
* Top Countries (Origin, Destination)
* Top Routes
* Interesting Aircraft (Miiltary, Government, Police, Civilian)
* Fastest Aircraft
* Slowest Aircraft
* Highest Aircraft
* Lowest Aircraft

## Setup

There are environment variables (`LATITUDE`, `LONGITUDE`, `RADIUS`) that can be used to only process aircraft data that falls within a particular boundary - similar to [planefence](https://github.com/sdr-enthusiasts/docker-planefence). Alternatively, setting the `RADIUS` to something larger than that of your SDR will mean all aircraft data is processed.

### Running in Docker (recommended)
* Clone this repository
* Populate `docker-compose.yml` with all required values
* Run `docker compose up -d --build`

### Running locally (eg. to develop)
* Clone this repository
* Create the postgres db (e.g. in a Docker container) - `/scripts/schema.sql` can be used to initialise the database
* Rename `.env-example` to `.env` and populate
* Change to the `core` folder e.g. `cd core`
* Compile with `go build -o skystats-daemon`
* Run the app `./skystats-daemon`
* It can be terminated via `kill $(cat skystats/core/skystats.pid)`

### Environment Variables

#### These environment variables should all be configured based on your setup, location etc

| Environment Variable | Description | Example |
|---|---|---|
| ADSB_HOST | IP of the ADSB receiver running readsb. | `192.168.1.50` |
| ADSB_PORT | Port on ADSB receiver where readsb aircraft.json being served. You can check via accessing `{ADSB_HOST}:{ADSB_PORT}/data/aircraft.json` in a browser and seeing if you see aircraft.json contents. | `8080` |
| POSTGRES_USER | Username you want to set for the postgresql database. | `admin` |
| POSTGRES_PASSWORD | Password you want to set for the postgresql database. | `password` |
| DB_USER | * Username of the postgresql database * Must match `POSTGRES_USER` * There are two places in the docker-compose.yml where it must be set;`app` and `db-init` | `admin` |
| DB_PASSWORD | * Password of the postgresql database * Must match `POSTGRES_PASSWORD` * There are two places in the docker-compose.yml where it must be set;`app` and `db-init` | `password` |
| LATITUDE | Lattitude of your receiver. | `XX.XXXXXX` |
| LONGITUDE | Longitude of your receiver. | `YY.YYYYYY` |
| RADIUS | Distance in km from your receiver that you want to record aircraft. Set to a distance greater than that of your receiver to capture all aircraft. | `1000` |
| DOMESTIC_COUNTRY | ISO 2-letter country code of the country your receiver is in - used to generate the "Domestic Airport" stats. | `GB` |

#### These environment variables most likely shouldn't be changed - keep as their default values

| Environment Variable | Description | Default Value |
|---|---|---|
| DOCKER_ENV | Flag used to skip daemonising the app when running in Docker. Do not change to `false` unless running outside Docker .(e.g. for development / debugging). | `true` |
| POSTGRES_DB | Name of the postgresql database. | `skystats_db` |
| DB_HOST | Name of the docker container hosting the postgresql database. | `skystats-db-docker` |
| DB_PORT | Default port for postgresql. | `5432` |
| DB_NAME | Name of the postgresql database. | `skystats_db` |
| API_PORT | Port that the Skystats API will be served on. | `8080` |
| ABOVE_RADIUS | Radius for the "Above Timeline" - currently only 20km supported. | `20` |
| ADSB_DB_AIRCRAFT_ENDPOINT | ADSB DB aircraft endpoint, used for registration data. | `https://api.adsbdb.com/v0/aircraft/` |
| ADSB_DB_CALLSIGN_ENDPOINT | ADSB DB callsign endpoint, used for route data. | `https://api.adsbdb.com/v0/callsign/` |
<!-- | PLANE_DB_URL | URL of the plane-alert-db list of interesting planes. | [Link](https://raw.githubusercontent.com/sdr-enthusiasts/plane-alert-db/refs/heads/main/plane-alert-db.csv) |
| IMAGE_DB_URL | URL of the plane-alert-db image urls. | [Link](https://raw.githubusercontent.com/sdr-enthusiasts/plane-alert-db/refs/heads/main/plane_images.csv) | -->


## Screenshots

### Home
![Home](docs/screenshots/1_Home.png)
</br>
![AboveMeModal](docs/screenshots/2_AboveMeModal.png)
</br>

### Route Stats
![RouteStats](docs/screenshots/3_RouteStats.png)
</br>

### Interesting Aircraft
![InterestingSeen](docs/screenshots/4_InterestingStats.png)
</br>
![InterestingModal](docs/screenshots/5_InterestingModal.png)
</br>

### Motion Stats
![MotionStats](docs/screenshots/6_MotionStats.png)


