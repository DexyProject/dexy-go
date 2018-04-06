# Markets

## Get /markets

Returns information for list of ```markets```

### Parameters
* **tokens [array]** ```required``` A json encoded array of token addresses. 
```json
["0xbebb2325ef529e4622761498f1f796d262100768"]
```

### Response Body
```json
[
  {
    "token": "0xbebb2325ef529e4622761498f1f796d262100768",
    "bid": 0.003,
    "ask": 0.004436923082864458,
    "volume": 0.01,
    "last": 0.0036923077377514803
  }
]
```
