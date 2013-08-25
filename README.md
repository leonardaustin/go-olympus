# Go Olympus
*Olympus was not shaken by winds nor ever wet with rain, nor did snow fall upon it, but the air is outspread clear and cloudless, and over it hovered a radiant whiteness. Here the major gods live and hold court.*

Olympus manages all services and routes requests to the correct location (both internally and externally). It authenticates tokens (if the service requires) and could be thought as the front controller (to all our services). Externally is can handle http json requests but internally it is all protobuf.

## Work Flow
### External
* Liten to external http requests and convert json to proto
* Authenticate request if required
* Add request to queue for correct service and listen for response
* Convert response to json and return

### Internal
* Allow services to register endpoints

## Server Stack
- Go
- protobuf
- Cassandra
- NSQ

## Client Stack
- Backbone.js
- Require.js
- jQuery ()
- Twitter Bootstrap (HTML/CSS/JS)