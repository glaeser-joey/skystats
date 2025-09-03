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

### Running in Docker (recommended)
* Clone this repository
* Copy the contents of `.env.example` into a new file called `.env`
* Populate `.env` with all required values. See [Environment Variables](#environment-variables)
* Run `docker compose up -d --build`
* The interface should be available on `localhost:5173` where localhost is the IP of the docker host

### Running locally (e.g. to develop)
* Clone this repository
* Create the postgres database (e.g. in a Docker container) - `/scripts/schema.sql` can be used to initialise the database
* Copy the contents of `.env.example` into a new file called `.env`
* Populate `.env` with all required values. See [Environment Variables](#environment-variables)
* Change to the `core` folder e.g. `cd core`
* Compile with `go build -o skystats-daemon`
* Run the app `./skystats-daemon`
    * It can be terminated via `kill $(cat skystats/core/skystats.pid)`
* Run the webserver 
    * Change to the /web directory e.g. `cd ../web`
    * Start the webserver with `npm run dev -- --host`

### Environment Variables

| Environment Variable | Description | Example |
|---|---|---|
| READSB_AIRCRAFT_JSON | URL of where readsb [aircraft.json](https://github.com/wiedehopf/readsb-githist/blob/dev/README-json.md) is being served e.g. http://yourhost:yourport/data/aircraft.json | `http://192.168.1.100:8080/data/aircraft.json` |
| DB_HOST | Postgres host. If running in docker this should be the name of the postgres container. If running locally it should be the IP/hostname of wherever postgres is hosted. | Docker: `skystats-db` <br/> Local: `192.168.1.10` |
| DB_PORT | Postgres port | `5432` |
| DB_USER | Postgres username | `user` |
| DB_PASSWORD | Postgres password | `1234` |
| DB_NAME | Postgres database name | `skystats_db` |
| DOMESTIC_COUNTRY_ISO | ISO 2-letter country code of the country your receiver is in - used to generate the "Domestic Airport" stats. | `GB` |
| LAT | Lattitude of your receiver. | `XX.XXXXXX` |
| LON | Longitude of your receiver. | `YY.YYYYYY` |
| RADIUS | Distance in km from your receiver that you want to record aircraft. Set to a distance greater than that of your receiver to capture all aircraft. | `1000` |
| ABOVE_RADIUS | Radius for the "Above Timeline" <br/> **Note: currently only 20km supported.** | `20` |

<br/>

## Advanced Use Cases

### Custom plane-alert-db csv

If you live in an area where you frequently see planes that you are not interested in, you can provide a custom version of [plane-alert-db](https://github.com/sdr-enthusiasts/plane-alert-db).

Add the following the the `.env` file:
```
PLANE_DB_URL=some/custom/location/plane-alert-db.csv
IMAGE_DB_URL=some/custom/location/plane_images.csv
```

And the following to `docker-compose.yml` under the `skystats-db-init` service:
```
- PLANE_DB_URL=${PLANE_DB_URL}
- IMAGE_DB_URL=${IMAGE_DB_URL}
```

**⚠️ The format of the csv must match the format of plane-alert-db**

<br/>

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