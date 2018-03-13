# Ticks

## Get /ticks

Returns tickers for a specified ```token```

### Parameters
* **token [string]** ```required``` - Returns all ticks where ```quote``` equals to ```token```.

```json
[
  {
    "pair": {
      "base": "0x0000000000000000000000000000000000000000",
      "quote": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"
    },
    "block": 4862998,
    "volume": "100",
    "open": "0.1",
    "close": "0.3",
    "high": "0.39",
    "low": "0.09",
    "timestamp": "1515233752"
  }
]
```
