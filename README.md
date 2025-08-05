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


### Running locally (eg. to develop)
* Clone this repository
* Create the postgres db (e.g. in a Docker container) - `/scripts/schema.sql` can be used to initialise the database
* Rename `.env-example` to `.env` and populate
* Change to the `core` folder e.g. `cd core`
* Compile with `go build -o skystats-daemon`
* Run the app `./skystats-daemon`
* It can be terminated via `kill $(cat skystats/core/skystats.pid)`

### Running in Docker
* Clone this repository
* Populate `docker-compose.yml` with all required values
* Run `docker compose up -d --build`


### Environment Variables

Either set in `.env` for local development, or `docker-compose.yml` when running in docker.

| DB_NAME                   | Name of the postgres database - `skystats_db`                                                                                                             |
|---------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| DB_USER                   | Postgres username - `admin`                                                                                                                               |
| DB_PASSWORD               | Postgres password - `password`                                                                                                                            |
| DB_HOST                   | Postgres host - `192.168.1.xxx`                                                                                                                           |
| DB_PORT                   | Postgres port - `5432`                                                                                                                                    |
| ADSB_HOST                 | IP of the ADSB receiver running readsb - `192.168.1.xxx`                                                                                                  |
| ADSB_PORT                 | Port on ADSB receiver where readsb aircraft.json being served - `8080` (Note: Not the `30047/tcp` port which streams JSON data)                           |
| LATITUDE                  | Lattitude of your receiver - `xx.xxxxxx`                                                                                                                  |
| LONGITUDE                 | Longitude of your receiver - `yy.yyyyyy`                                                                                                                  |
| RADIUS                    | Distance in km from your receiver that you want to record aircraft. Set to a distance greater than that of your receiver to capture all aircraft - `1000` |
| ADSB_DB_AIRCRAFT_ENDPOINT | ADSB DB aircraft endpoint, used for registration data - `https://api.adsbdb.com/v0/aircraft/`                                                             |
| ADSB_DB_CALLSIGN_ENDPOINT | ADSB DB callsign endpoint, used for route data - `https://api.adsbdb.com/v0/callsign/`                                                                    |

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


