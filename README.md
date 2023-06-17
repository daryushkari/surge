Surge service:

## Project Description:

surge service is a microservice that controls price coefficients by districts.
it divides Tehran city to 22 municipal districts and controls price by demand in each district.
each time a customer requests a ride, surge service will be called by rest API that has latitude and longitude
of origin of surge service. then it will get district polygons from open street map APIs and check that request
belongs to which district and saves each request by time and district id in redis cache.
also it has a cronjob that runs every several minutes and removes old requests by district id that is no longer needed.
It has request thresholds and when requests in predefined dynamic time window surpasses the defined threshold,
we notify pricing service by nats and if pricing service acknowledge,
surge service will save last coefficient that sent pricing service and will not communicate pricing service
until coefficient changes for that specific district.

## Setup:
to run project you only need to run 2 commands:
```
   make config
   docker-compose up
```
also make sure that config settings is correct for your environment.
service APIs are defined in postman directory.

# Architecture:

project has 3 main layers:
- delivery
- use-case
- repository

delivery:
delivery layer defines APIs and handler functions and then when request calls it validates incoming 
requests, clean them and then call use-case layer to process request.

use-case:
in use-case layer we handle business logic and do whatever logic handling we need. we call repository layer
to get data we need and make decisions based on request and data we have previously stored.

repository:
in repository layer we define functions to insert, update, delete or read from database or define more complicated
queries to run on database.

also our project has other main parts including:
- config
- cronjob
- entity
- pkg

config:
we read configs from config file in startup.

cronjob:
define periodic jobs that should be run

entity:
define data entities we have 2 directories in it. domain saves data that should be stored in database
and request model that defines API request response models.

pkg:
we use basic functions or wrappers that is independent of other parts of project and consists connecting to
tools like database and ... or calling external APIs

