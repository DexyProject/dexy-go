# History

## Get /trades

Returns trades for a specified ```token```

### Parameters
* **token [string]** ```required``` - Returns all trades where ```token``` in either side of the trade is equal to the address. 
* **limit [int]** ```Default: 100``` - Amount of trades. 
* **user [maker]** ```optional``` - Returns all trades where ```maker``` equals to specified address

```json
[
  {
    "tx": "0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
    "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
    "block": 4862998,
    "timestamp": "1515233752",
    "taker": "0x997919a608788621dd48b3896f78dcda682fe91d",
    "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
    "take": {
      "token": "0x0000000000000000000000000000000000000000",
      "amount": "3000000000000000000"
    },
    "make": {
      "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
      "amount": "300000000000000000000"
    }
  }
]
```