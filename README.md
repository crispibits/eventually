# Eventually

An experiment in writing an events manager

Initial loads of events can be done with json:
in `objects/events.json` add events to an array like so:
```json
[
    {
        "title": "Overthrowing The World With A Teaspoon",
        "content": "A frankly ludicrous talk about how to overthrow the worl with a teaspoon",
        "costs": [
            { "ticketType": "standard",
               "amount": 7.00 },
            { "ticketType": "concession",
               "amount" : 5.00 }
        ],
        "times": [
            {
                "Doors" : "2025-07-15T19:00",
                "Start" : "2025-07-15T19:30",
                "End" : "2025-07-15T22:00"
            }
        ],
        "location": "MAIN_AUDITORIUM",
        "owner": "Crispin",
        "tags": [
            "provisional",
            "art"
        ],
        "ticketUrls": [
            {
                "provider": "ticketyboo",
                "url": "https://www.example.com/teaspoon"
            }
        ],
        "restrictions": [
            "Over 18's"
        ]
    }
]
```

In the same way add locations to `objects/locations.json`:
```json
[
    {
        "name": "MAIN_AUDITORIUM",
        "displayName": "The Main Auditorium",
        "address": "Teaspoon Avenue, North Southamptonshire",
        "longitude": "7.4",
        "latitude": "9.2",
        "url": "https://example.com"
    }
]
```
