# media-lookup

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

Required flags: 
    - query
Optional flags: 
    - nsfw
        - Defaults to false
    - language 
        - Defaults to 'en'
    - page
        - Defaults to 1
