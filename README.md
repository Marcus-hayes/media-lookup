# Media lookup
## Overview
A Go-Cobra CLI tool for searching and getting metadata on movies, TV shows, and famous persons

## Required Environment Variables
* TMDB_API_Key
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

#### Command Flags
Flags:
| Flag Name | Type      | Description                |
| :-------- | :-------- | :-------------------------    |
| `query`   | `string`  | **Required** Query string |
| `nsfw`    | `bool`    | Optional NSFW filter; default is false |
| `language` | `string` | Optional language filter; default is 'en' |
| `page`    | `int32`   | Optional page filter; default is 1 |

### Get Movie, TV Show, or Person Metadata
`./media-lookup details --id "12345" --media-type "[person|tv|movie]"`

#### Command Flags
Flags:
| Flag Name | Type      | Description                |
| :-------- | :-------- | :-------------------------    |
| `id`   | `string`  | **Required** ID string of TMDB object |
| `media-type`    | `string`    | **Required** Media type of TMDB Object {"person", "tv", "movie"} |