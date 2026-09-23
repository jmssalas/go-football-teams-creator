# Football Teams Creator

A small web application for managing players, seasons, match results, and generating balanced football teams from historical performance.

The project combines a Go backend with a lightweight browser frontend. It tracks player statistics from matches, keeps seasons separate, and exposes an API consumed by the frontend.

## Features

- Manage players
- Create and view seasons
- Record match results with team assignments and scorelines
- Calculate per-player statistics such as wins, draws, losses, goals for/against, points, and win percentage
- Generate balanced teams from player statistics
- Save generated teams for later use
- Serve the frontend directly from the Go application

## Tech Stack

- Go
- Gin
- SQLite
- GORM
- HTML, CSS, and JavaScript

## How It Works

The backend persists data in SQLite and exposes a REST API used by the frontend.

The main resources are:

- **Players** — player information and season participation
- **Seasons** — separate periods for tracking historical data
- **Matches** — completed matches, including scorelines and team assignments
- **PlayerMatch** — relationship between players and matches, including which team they played for

Player statistics are derived from match results rather than stored as independent cumulative values. This keeps the data consistent and allows statistics to be calculated for the selected season.

The team generation algorithm uses backtracking to find balanced combinations of players based on their historical performance.

## Running the App

### Prerequisites

- Go 1.27+

### Configuration

The application requires the following environment variables:

```bash
PORT=8080
SQLITE_DATABASE_PATH=/absolute/path/to/football-teams-creator.db
TEAMS_DATA_PATH=/absolute/path/to/teams.json
```

### Start the application

From the project root:

```bash
go mod download

PORT=8080 \
SQLITE_DATABASE_PATH="$(pwd)/football-teams-creator.db" \
TEAMS_DATA_PATH="$(pwd)/data/teams.json" \
go run .
```

Then open:

```text
http://localhost:8080
```

The Go application serves the frontend assets directly from the `web/` directory, so no separate frontend build step is required.

## Deployment

The application is packaged as a Docker image and published to GitHub Container Registry on every push to `master`.

A public demo instance is deployed as a Docker container using Docker Compose. Runtime configuration is provided through environment variables, while application data is persisted outside the container.

The demo instance is available at:

https://demo-ftc.jmssalas.com/

## API

The application exposes a REST API for the main resources.

Some of the available endpoints are:

| Method   | Endpoint                           | Description                                                                                  |
| -------- | ---------------------------------- | -------------------------------------------------------------------------------------------- |
| `POST`   | `/api/players`                     | Create a player                                                                              |
| `GET`    | `/api/players?seasonId=<seasonId>` | List players with their statistics for the selected season, or the current season if omitted |
| `DELETE` | `/api/players/:id`                 | Remove a player                                                                              |
| `POST`   | `/api/seasons`                     | Create a season                                                                              |
| `GET`    | `/api/seasons`                     | List seasons                                                                                 |
| `GET`    | `/api/seasons/current`             | Get the current season                                                                       |
| `POST`   | `/api/matches`                     | Record a completed match                                                                     |
| `GET`    | `/api/teams`                       | Fetch saved teams                                                                            |
| `POST`   | `/api/teams`                       | Store generated teams                                                                        |

## Engineering Notebook

This project was rebuilt after the original version had been used in practice and its structure became increasingly difficult to extend.

The rebuild was not about replacing everything. Parts of the original implementation worked perfectly well and were deliberately kept.

**Read the full story:** [Building It Twice](https://jmssalas.com/notebook/building-it-twice)
