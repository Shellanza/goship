# GSW - Goship Shipreporting Webservice 

Note: calls marked w/ [*MP] are marked for porting. Those w/ [*OK] are completed, deployed and available for production. Calls w/ [*SB] are in stand-by, candidated to be rejected.

__Version: v0.2.5__

## List of verified & Active services:

- [A.2 - Ships at roadstead](#a2-at-roadstead-ok)

- [A.3 - Ships at mooring](#a3-moored-ok)

- [A.4 - Ships arrived today](#a4-arrivals-ok)

- [A.5 - Ships departed today](#a5-departures-ok)

- [A.6 - Arrivals previsions today](#a6-arrival-previsions-ok)

- [A.9 - Shifting previsions today](#a9-shifting-previsions-ok)

- [A.10 - Departure previsions today](#a10-departure-previsions-ok)

Register data:

- [C.1 - Arrivals](#c1-arrivals-ok)
  
- [C.2 - Moored](#c2-moored-ok)

- [C.3 - At roadstead](#c3-roadstead-ok)

- [C.4 - Departures](#c4-departures-ok)

Meteo data:

- [E.2 - Active stations](#e2-active-meteo-stations)
  
  _________

## A - LIVE DATA SERVICES:

# A1. __Active trips__: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/activeTripsNow?id_portinformer=<id_portinformer>
```

# A2. __At roadstead__: [*OK]

Request:

```bash
Request:
http://<REMOTE_IP>:8000/anchored/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, anchoringTime, currentActivity, anchoragePoint, shipType, iso3, grossTonnage, length, width, agency, shippedGoods]

res := map[string]string{
    "id_trip":          idControlUnitDataStr.String,
    "ship":             shipName.String,
    "ship_type":        shipType.String,
    "anchoring_time":   anchoringTime.String,
    "current_activity": currentActivity.String,
    "anchorage_point":  anchoragePoint.String,
    "iso3":             iso3.String,
    "gross_tonnage":    grossTonnage.String,
    "length":           length.String,
    "width":            width.String,
    "agency":           agency.String,
    "shipped_goods":    shipped_goods.String,
    }

```

# A3. __Moored__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/moored/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, mooringTime, currentActivity, quay, shippedGoods, country, grossTonnage, length, width,  agency]

res := map[string]string{
    "id_trip":          idControlUnitDataStr.String,
    "ship":             shipName.String,
    "mooring_time":     mooringTime.String,
    "current_activity": currentActivity.String,
    "quay":             quay.String,
    "shipped_goods":    shippedGoods.String,
    "iso3":             iso3.String,
    "gross_tonnage":    grossTonnage.String,
    "ships_length":     length.String,
    "ships_width":      width.String,
    "ship_type":        shipType.String,
    "agency":           agency.String,
    }

```

# A4. __Arrivals__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, country, width, length, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idTrip.String,
    "ship_name":              shipName.String,
    "ship_type":              shipType.String,
    "ts_sighting":            tsSighting.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "port_destination":       portDestination.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    }


```

# A5. __Departures__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/departures/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, portDestination]


res := map[string]string{
    "id_trip":           idTrip.String,
    "ship_name":         shipName.String,
    "ship_type":         shipType.String,
    "ts_out_of_sight":   tsOutOfSight.String,
    "ship_flag":         shipFlag.String,
    "ship_width":        shipWidth.String,
    "ship_length":       shipLength.String,
    "gross_tonnage":     grossTonnage.String,
    "net_tonnage":       netTonnage.String,
    "draft_aft":         draftAft.String,
    "draft_fwd":         draftFwd.String,
    "agency":            agency.String,
    "last_port_of_call": lastPortOfCall.String,
    "port_destination":  portDestination.String,
    }
```

# A6. __Arrival previsions__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/arrivalPrevisions/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idControlUnitDataStr.String,
    "ship":                   shipName.String,
    "ts_arrival_prevision":   tsArrivalPrevision.String,
    "ship_type":              shipType.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    "cargo_on_board":         cargoOnBoard.String,
    }

```

# A7. __Shipped goods__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, quantity, unit, goodsCategory, shipType, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, groupCategory, macroCategory]

res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "quantity":       quantity.String,
    "unit":           unit.String,
    "goods_category": goodsCategory.String,
    "ship_type":      shipType.String,
    "ship_flag":      shipFlag.String,
    "ship_width":     shipWidth.String,
    "shipLength":     shipLength.String,
    "grossTonnage":   grossTonnage.String,
    "netTonnage":     netTonnage.String,
    "groupCategory":  groupCategory.String,
    "macroCategory":  macroCategory.String,
    }

```

# A8. __RO/RO + RO/PAX__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/trafficListToday/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, numContainer, numPassengers, numCamion, numFurgoni, numRimorchi, numAuto, numMoto, numCamper, tons, numBus, numMinibus, trafficListMvntType, trafficListCategories, quay]


res := map[string]string{
    "id_trip":        idTrip.String,
    "ship_name":      shipName.String,
    "num_container":  numContainer.String,
    "num_passengers": numPassengers.String,
    "num_camion":     numCamion.String,
    "num_furgoni":    numFurgoni.String,
    "num_rimorchi":   numRimorchi.String,
    "num_auto":       numAuto.String,
    "num_moto":       numMoto.String,
    "num_camper":     numCamper.String,
    "tons":           tons.String,
    "num_bus":        numBus.String,
    "num_minibus":    numMinibus.String,
    "mvnt_type":      mvntType.String,
    "description":    description.String,
    "quay":           quay.String,
    }


```

# A9. __Shifting previsions__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/shiftingPrevisions/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, ts_shifting_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage,net_tonnage, draft_aft, draft_fwd, agency, starting_quay_berth, starting_roadstead, stop_quay_berth, stop_roadstead, cargo_on_board]


res := map[string]string{
    "ship":                 ship.String,
    "tsDeparturePrevision": tsShiftingPrevision.String,
    "shipType":             shipType.String,
    "shipFlag":             shipFlag.String,
    "shipWidth":            shipWidth.String,
    "shipLength":           shipLength.String,
    "grossTonnage":         grossTonnage.String,
    "netTonnage":           netTonnage.String,
    "draftAft":             draftAft.String,
    "draftFwd":             draftFwd.String,
    "agency":               agency.String,
    "destinationPort":      destinationPort.String,
    "startingQuayBerth":    startingQuayBerth.String,
    "startingRoadstead":    stopRoadstead.String,
    "stopQuayBerth":        stopQuayBerth.String,
    "stopRoadstead":        stopRoadstead.String,
    "cargoOnBoard":         cargoOnBoard.String,
    }


```

# A10. __Departure previsions__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/departurePrevisions/<id_portinformer>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, ts_departure_prevision, ship_type, ship_flag, ship_width, ship_length, gross_tonnage, net_tonnage, draft_aft, draft_fwd, agency, destination_port, starting_quay_berth, starting_roadstead, cargo_on_board]

res := map[string]string{
    "ship":                   ship.String,
    "ts_departure_prevision": tsDeparturePrevision.String,
    "ship_type":              shipType.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "destination_port":       destinationPort.String,
    "starting_quay_berth":    startingQuayBerth.String,
    "starting_roadstead":     startingRoadstead.String,
    "cargo_on_board":         cargoOnBoard.String,
    }

```

## B - ARCHIVE DATA SERVICES:

# B1. __Trips archive [global recap, one row per trip]__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/tripsArchive?id_portinformer=<ID_PORTINFORMER>
```

# B2. __Trips archive [global recap, one row per commercial operation]__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/tripsArchiveMultiRows?id_portinformer=<ID_PORTINFORMER>
```

# B3. __Trip data archive__ [shipreport core]: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/shipReportList?id_portinformer=<ID_PORTINFORMER>
```

# B4. __Trip data archive detailed__ [shipreport]: [*SB]

Request:

```bash
http://<REMOTE_IP>:8000/shipReportDetails?id_portinformer=<ID_PORTINFORMER>
```

# B5. __Arrivals archive__: [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsArchive?id_portinformer=<id_portinformer>
```

# B6. __Departures archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/departuresArchive?id_portinformer=<id_portinformer>
```

# B7. __Shipped goods archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsArchive?id_portinformer=<id_portinformer>
```

# B8. __Traffic list archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/trafficListArchive?id_portinformer=<id_portinformer>
```

## C - DAILY REGISTER SERVICES:

# C1. __Arrivals:__ [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/arrivalsRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [ship, tsArrivalPrevision, shipType, country, width, length, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, destinationQuayBerth, destinationRoadstead, cargoOnBoard]

res := map[string]string{
    "id_trip":                idTrip.String,
    "ship_name":              shipName.String,
    "ship_type":              shipType.String,
    "ts_sighting":            tsSighting.String,
    "ship_flag":              shipFlag.String,
    "ship_width":             shipWidth.String,
    "ship_length":            shipLength.String,
    "gross_tonnage":          grossTonnage.String,
    "net_tonnage":            netTonnage.String,
    "draft_aft":              draftAft.String,
    "draft_fwd":              draftFwd.String,
    "agency":                 agency.String,
    "last_port_of_call":      lastPortOfCall.String,
    "port_destination":       portDestination.String,
    "destination_quay_berth": destinationQuayBerth.String,
    "destination_roadstead":  destinationRoadstead.String,
    }


```

# C2. __Moored:__ [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/mooredRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsMooring, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, agency]

res := map[string]string{
    "id_trip":       idTrip.String,
    "ship_name":     shipName.String,
    "ship_type":     shipType.String,
    "ts_mooring":    tsMooring.String,
    "ship_flag":     shipFlag.String,
    "ship_width":    shipWidth.String,
    "ship_length":   shipLength.String,
    "gross_tonnage": grossTonnage.String,
    "net_tonnage":   netTonnage.String,
    "agency":        agency.String,
    }
```

# C3. __Roadstead:__ [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/roadsteadRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsAnchoring, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, agency]

res := map[string]string{
    "id_trip":       idTrip.String,
    "ship_name":     shipName.String,
    "ship_type":     shipType.String,
    "ts_anchoring":  tsAnchoring.String,
    "ship_flag":     shipFlag.String,
    "ship_width":    shipWidth.String,
    "ship_length":   shipLength.String,
    "gross_tonnage": grossTonnage.String,
    "net_tonnage":   netTonnage.String,
    "agency":        agency.String,
    }
```

# C4. __Departures:__ [*OK]

Request:

```bash
http://<REMOTE_IP>:8000/departuresRegister/<ID_PORTINFORMER>/<TIMESTAMP_START>/<TIMESTAMP_STOP>
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idTrip, shipName, shipType, tsOutOfSight, shipFlag, shipWidth, shipLength, grossTonnage, netTonnage, draftAft, draftFwd, agency, LPC, portDestination]


res := map[string]string{
    "id_trip":           idTrip.String,
    "ship_name":         shipName.String,
    "ship_type":         shipType.String,
    "ts_out_of_sight":   tsOutOfSight.String,
    "ship_flag":         shipFlag.String,
    "ship_width":        shipWidth.String,
    "ship_length":       shipLength.String,
    "gross_tonnage":     grossTonnage.String,
    "net_tonnage":       netTonnage.String,
    "draft_aft":         draftAft.String,
    "draft_fwd":         draftFwd.String,
    "agency":            agency.String,
    "last_port_of_call": lastPortOfCall.String,
    "port_destination":  portDestination.String,
    }
```

# C5. __Shiftings:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/registerShiftings?id_portinformer=<ID_PORTINFORMER>
```

# C6. __Arrival previsions:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/registerPlannedArrivals?id_portinformer=<ID_PORTINFORMER>
```

# C7. __Shipped goods:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/registerShippedGoods?id_portinformer=<ID_PORTINFORMER>
```

# C8. __RO/RO + RO/PAX:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/registerTrafficList?id_portinformer=<ID_PORTINFORMER>
```

## D - BUSINESS INTELLIGENCE SERVICES: ##

# D1. __Shiftings/maneuverings [per quay/berth]:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/tripsManeuverings?id_portinformer=<ID_PORTINFORMER>
```

# D2. __Shipped goods recap:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/shippedGoodsRecap?id_portinformer=<ID_PORTINFORMER>
```

# D3. __RO/RO + RO/PAX recap:__ [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/trafficListRecap?id_portinformer=<ID_PORTINFORMER>
```

## E - METEO DATA ##
# E1. __Meteo data archive__: [*MP]

Request:

```bash
http://<REMOTE_IP>:8000/meteoArchive?id_portinformer=<ID_PORTINFORMER>
```

# E2. __Active meteo stations__:

Request:

```bash
http://<REMOTE_IP>:8000/weatherActiveStations
```

Response:

```bash
# Content-Type: application/json; charset=UTF-8

# data := [idPortinformer, portinformerCode, tsFirstCreated, isActive]

res := map[string]string{
    "id_portinformer":   idPortinformer.String,
    "portinformer_code": portinformerCode.String,
    "ts_first_created":  tsFirstCreated.String,
    "is_active":         isActive.String,
}
```

# Install Go and (first) run 
ref. https://golang.org/doc/install?download=go1.12.10.linux-386.tar.gz

```bash
$ go version
go version go1.12.10 linux/amd64

$ export PATH=${PATH}:/usr/local/go/bin/ # FIX w/ your installation path

# Pass vars at runtime or add to .bash_profile
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}

# Get goship and deploy
$ go get github.com/deeper-x/goship

$ cd ${GOPATH}/src/github.com/deeper-x/goship
$ cat > .env <<HEREDOC 
DB_DSN="postgres://<user>:<passwd>@127.0.0.1/<db>"
HEREDOC

$ go get -d ./...    
$ go install src/main.go 
$ goship # RUN FOR TEST
Now listening on: http://localhost:8000
Application started. Press CTRL+C to shut down.

# Close test instance
$ <CTRL+C>  

```

# Integration Test (httpexpect)
```bash
$ go test -v ./...
=== RUN   TestMain
--- PASS: TestMain (0.44s)
PASS
ok  	github.com/deeper-x/goship/src	(cached)
```

# Systemd configuration

```bash
# Service install instructions: 
# This is a service file template you can start from 
$ sudo cat > /usr/lib/systemd/system/goship.service <<HEREDOC

[Unit]
Description=Shipreporting service middleware
Documentation=https://github.com/deeper-x/goship
After=network.target

[Service]
Type=simple
User=<YOUR_USER>
WorkingDirectory=<YOUR_GOPATH>/bin/
ExecStart=<YOUR_GOBIN>/goship
Restart=on-failure

[Install]
WantedBy=multi-user.target

HEREDOC
```


