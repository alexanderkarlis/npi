
# NPI 
This repository is a wrapper for [CMS's api](https://npiregistry.cms.hhs.gov/). Check out NPPES docs for
API usage. Almost an exact port in Go, for all exposed queryables.
## Run docker container
```shell
> docker build -t npi .
> docker run -p 8888:8080 npi --name npiregistry 
```
## Example
```shell
> curl -d '{"state": "ny", "city": "cranford"}' -X POST 'http://127.0.0.1:8888/npi'
> curl -d '{"number": "1689872012"}' -X POST 'http://localhost:8888/npi'
"result_count": 1,
"results": [
  {
    "addresses": [
      {
        "address_1": "68 HARRIS BUSHVILLE RD",
        "address_2": "",
        "address_purpose": "LOCATION",
        "address_type": "DOM",
        "city": "MONTICELLO",
        "country_code": "US",
        "country_name": "United States",
        "fax_number": "845-796-1404",
        "postal_code": "127013027",
        "state": "NY",
        "telephone_number": "845-794-0996"
      },
      {
        "address_1": "191 N LEHIGH AVE",
        "address_2": "",
        "address_purpose": "MAILING",
        "address_type": "DOM",
        "city": "CRANFORD",
        "country_code": "US",
        "country_name": "United States",
        "fax_number": "908-272-0212",
        "postal_code": "070163040",
        "state": "NJ",
        "telephone_number": "908-644-5182"
      }
    ],
    "basic": {
      "credential": "M.D.",
      "enumeration_date": "2007-07-06",
      "first_name": "BRIJ",
      "gender": "M",
      "last_name": "MOHAN",
      "last_updated": "2010-03-23",
      "name": "MOHAN BRIJ",
      "sole_proprietor": "NO",
      "status": "A"
    },
    "created_epoch": 1183680000,
    "enumeration_type": "NPI-1",
    "identifiers": [
      {
        "code": "01",
        "desc": "Other",
        "identifier": "255216",
        "issuer": "NY LICENSE",
        "state": "NY"
      }
    ],
    "last_updated_epoch": 1269302400,
    "number": 1689872012,
    "other_names": [],
    "taxonomies": [
      {
        "code": "208600000X",
        "desc": "Surgery",
        "license": "R8004",
        "primary": false,
        "state": "IA"
      },
      {
        "code": "208600000X",
        "desc": "Surgery",
        "license": "MD438480",
        "primary": false,
        "state": "PA"
      },
      {
        "code": "208600000X",
        "desc": "Surgery",
        "license": "255216-1",
        "primary": true,
        "state": "NY"
      }
    ]
  }
]
}⏎       

