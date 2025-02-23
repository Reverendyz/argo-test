# TIMER API

Wrongly named by purpose. But the functionality is all that matters.

## Endpoints & Returns

```sh

# Method: GET
# Body: none
# Location: /
# Return:
# {
#   "message": "ok"
# }


# Method: POST
# Body: json object with username & password keys
# Location: /
# Return:
#
# {
#   "token": "<jwt_token>"
# }

# Authenticated Requests

#Method: GET
#Body: none
#Location: /timer
#Return:
#{
#  hostname: "<container hostname>",
#  date: "<Current date and time in 'YYYY-mm-ddTHH:mm:ss' format>"
#}
```
