# Football Teams Creator

A rebuilt version of a small application for creating balanced football teams based on player statistics.

This project started as a rewrite of an application that had already been used in practice by a local football club. The first version worked well, but as new requirements emerged, its internal structure became increasingly difficult to extend.

Rather than continuing to refactor the original project, I decided to rebuild it using what I had learned from real usage.

## How It Works

Players have a set of statistics based on their previous matches:

* Wins
* Draws
* Losses

These statistics are derived from the matches played rather than stored as cumulative values.

When creating teams, the application uses the players' scores to find a balanced distribution. The team generation algorithm is based on **backtracking**, which remains appropriate for the relatively small number of players involved.

## Technology

* Go
* SQLite

The frontend remains responsible for team generation, while the backend exposes the application through an API.

## Engineering Notebook

This project is part of my Engineering Notebook, where I document the engineering decisions, problems, and lessons behind the projects I build.

The rewrite was not about replacing everything. Parts of the original implementation worked perfectly well and were deliberately kept.

**Read the full story:** [Building It Twice](https://jmssalas.com/notebook/building-it-twice)
