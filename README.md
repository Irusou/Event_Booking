### Project Description
A GO-powered "Event Booking" REST API

### Supported Routes
- GET /events                   -> get a list of available events
- GET /events/<id>              -> get a single event
- POST /events                  -> create a bookable event          *Auth Required 
- PUT /events/<id>              -> update an event                  *Auth Required *Only by creator
- DELETE /events/<id>           -> delete an event                  *Auth Required *Only by creator
- POST /signup                  -> create a new user
- POST /login                   -> authenticate user                *Auth Token (JWT) 
- POST /events/<id>/register    -> register user for event          *Auth Required
- DELETE /events/<id>/register  -> cancel registration              *Auth Required
