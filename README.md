# Media lookup
## Overview
A Go-Cobra CLI tool for searching and getting metadata on movies, TV shows, and famous persons

## Required Environment Variables
* TMDB API Key
    - Can be obtained at following link:
        https://developer.themoviedb.org/reference/intro/authentication#api-key-quick-start
    - To set environment variables, I recommend this link for instruction:
        https://www3.ntu.edu.sg/home/ehchua/programming/howto/Environment_Variables.html

## Set-up
`go build .`

## Commands
### Get version
`./media-lookup version`

### Search for Movie, TV Show, or Person
`./media-lookup search --query "SpongeBob" --nsfw false --page 1 --language "en"`

## Search Flags
Required flags:
* query
    - string

Optional flags:
- nsfw
    - bool
    - Defaults to false
- language
    - string
    - Defaults to 'en'
- page
    - int
    - Defaults to 1

### Get Movie, TV Show, or Person Metadata
`./media-lookup details --id "12345" --media-type "[person|tv|movie]"`

## Search Flags
Required flags:
* id
    - string
* media-type
    - string