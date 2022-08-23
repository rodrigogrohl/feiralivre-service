#!/bin/sh

curl -X GET -s localhost:8080/v1/streetmarket/ -d '{"id": 100}'

curl -X POST -s localhost:8080/v1/streetmarket/ -d '{"long": 657881, "lat": -657881, "sector": 5001, "area": 99, "dist_code": 458, "district": "MOEMA", "subtown_code": 546, "subtown": "VILA MARIANA", "region_5": "SUL", "region_8": "SUL B", "name": "FEIRA DO IBIRAPUERA", "registry": "1869-X", "addr": "AV IBIRAPUERA", "number": "1800", "neighborhood": "AV BRASIL", "reference": "PQ DO IBIRAPUERA"}' 

curl -X POST -s localhost:8080/v1/streetmarket/query/ -d '{"name": "IBIR"}'

curl -X POST -s localhost:8080/v1/streetmarket/query/ -d '{"name": "IBIR"}'