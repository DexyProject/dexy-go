# Markets

## Get /markets

Returns information for list of ```markets```

### Parameters
* **tokens [array]** ```required``` A json encoded array of token addresses. 
```json
["0xbebb2325ef529e4622761498f1f796d262100768", "0xe85b843180efce2de43aa6adcf04af7bf1f50c79"]
```

### Response Body
```json
{
    "0xbebb2325ef529e4622761498f1f796d262100768": {
        "bid": {
            "quote": "1000000",
            "base": "10000000"
        },
        "ask": {
            "quote": "1000000",
            "base": "10000000000000000"
        }
    },
    "0xe85b843180efce2de43aa6adcf04af7bf1f50c79": {
        "bid": {
            "quote": "1000000",
            "base": "10000000"
        },
        "ask": {
            "quote": "1000000",
            "base": "10000000000000000"
        }
    }
}
```
