# EMS / REST API Project

This is a small project that I do to quickly learn the basics of go / docker / adminer / echo / influxdb.
It consists of creating an EMS API to post values concerning a certain meter and a certain site, the values can then be read and visualized in Grafana or in InfluxDB UI.

## Run the project

In the first place, you have to change the settings in .env.example file to your liking.

Next, we can start up the EMS:
* Open the terminal in the project folder
* Use docker compose up --build (to start the services)
* In any browser, copy paste this link : **localhost:8080**

*It is the Adminer visualization to help you understand the relational database and every variable.*

You can now use the methods in Yaak to work with the EMS tool.

## List of endpoints

For the sites, they can have one of the next 4 types:
* Office
* Factory
* Warehouse
* Other

For the meters, they can have one of the next 3 types:
* Electricity
* Water
* Gas

*The unit must be specified*

To post data/values with Yaak, you must use the POST method on **localhost:8000/meters/:meter_id/readings**. In the JSON body you will need to specify wich **meter_id** you post the data, and what is the value. For example:

*{*
  **"meter_id": "d69345c0-0853-48b7-ac58-6497c5ef350c",*
  *"value": 144*
*}*

To read the data/values with Yaak, you must use the GET method on the same url: **localhost:8000/meters/:meter_id/readings**.

## Check the data (Grafana/InfluxDB UI)

The Grafana visualization is available on **localhost:3000**. In order to access the database:
* Base user and password are **admin** and **admin** respectivelly
* Go to Data sources under Connections
* Select InfluxDB type
* Select Flux query language
* The http url is: **http://influxdb:8086**
* Activate the basic auth and configure the details with watever you put in the .env.example file
* Same with the InfluxDB details
* Click **Save & test**

*You are good to go!*

The InfluxDB UI visualization is available on **localhost:8086**. In order to access the database:
* User and password are the same as in the .env.example file

*You are good to go!*