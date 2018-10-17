# Recommendation System

## Goal

Give a recommendation to the store owners/managers on the location for opening a new store.

## Current Implementation

We are using a government sourced dataset that contains each zip code in the country and a corresponding population density. We filter that data based on predifined zipcodes that belong to each county: San Jose, San Mateo and pick the largest from each.

## Next Steps

* Need to figure out an intelligent way for finding a more precise location for a new store rather than just a Zipcode.

## Running Instructions

Run in the terminal:

`go build logic.go`

Should printout a list of zip codes, country name and corresponding densities.