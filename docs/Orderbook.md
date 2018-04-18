# Orders

## Get /orderbook

Returns Asks and Bids for a specified ```token```

### Parameters
* **token [string]** ```required``` - Returns all orders where ```token``` in either side of the book is equal to the address. 
* **limit [int]** ```Default: 100``` - Amount of orders to return. 

```json
{
  "asks": 
  [
    {
      "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
      "take": {
        "token": "0x0000000000000000000000000000000000000000",
        "amount": "3000000000000000000"
      },
      "make": {
        "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
        "amount": "300000000000000000000"
      },
      "expires": 1514892553,
      "nonce": 12,
      "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
      "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",
      "signature": {
        "v": 12,
        "r": "0x9242685bf161793cc25603c231bc2f568eb630ea16aa137d2664ac8038825608",
        "s": "0x4f8ae3bd7535248d0bd448298cc2e2071e56992d0774dc340c368ae950852ada",
        "sig_mode": 0
      }
    }
  ],
  "bids": [
    {
      "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
      "take": {
        "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
        "amount": "300000000000000000000"
      },
      "make": {
        "token": "0x0000000000000000000000000000000000000000",
        "amount": "3000000000000000000"
      },
      "expires": 1514892553,
      "nonce": 12,
      "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
      "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",        
      "signature": {
        "v": 27,
        "r": "0x61a3ed31b43c8780e905a260a35faefcc527be7516aa11c0256729b5b351bc33",
        "s": "0x40349190569279751135161d22529dc25add4f6069af05be04cacbda2ace2254",
        "sig_mode": 0
      }
    }
  ]
}
```

## Get /orders

Returns orders ```token```

### Parameters
* **token [string]** ```required``` - Returns all orders where ```token``` in either side of the book is equal to the address. 
* **limit [int]** ```Default: 100``` - Amount of orders to return for both sides of the book. 
* **maker [string]** ```optional``` - Returns all orders where ```maker``` equals to specified address

```json
[
  {
    "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
    "take": {
      "token": "0x0000000000000000000000000000000000000000",
      "amount": "3000000000000000000"
    },
    "make": {
      "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
      "amount": "300000000000000000000"
    },
    "expires": 1514892553,
    "nonce": 12,
    "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
    "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",
    "signature": {
      "v": 12,
      "r": "0x9242685bf161793cc25603c231bc2f568eb630ea16aa137d2664ac8038825608",
      "s": "0x4f8ae3bd7535248d0bd448298cc2e2071e56992d0774dc340c368ae950852ada",
      "sig_mode": 0
    }
  },
  {
    "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
    "take": {
      "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
      "amount": "300000000000000000000"
    },
    "make": {
      "token": "0x0000000000000000000000000000000000000000",
      "amount": "3000000000000000000"
    },
    "expires": 1514892553,
    "nonce": 12,
    "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
    "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",
    "signature": {
      "v": 27,
      "r": "0x61a3ed31b43c8780e905a260a35faefcc527be7516aa11c0256729b5b351bc33",
      "s": "0x40349190569279751135161d22529dc25add4f6069af05be04cacbda2ace2254",
      "sig_mode": 0
    }
  }
]
```

## GET /orders/{order}

Returns order by the ```hash```

### Response Body

```json
{
  "hash": "0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
  "take": {
    "token": "0x0000000000000000000000000000000000000000",
    "amount": "3000000000000000000"
  },
  "make": {
    "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
    "amount": "300000000000000000000"
  },
  "expires": 1514892553,
  "nonce": 12,
  "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
  "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",
  "signature": {
    "v": 12,
    "r": "0x9242685bf161793cc25603c231bc2f568eb630ea16aa137d2664ac8038825608",
    "s": "0x4f8ae3bd7535248d0bd448298cc2e2071e56992d0774dc340c368ae950852ada",
    "sig_mode": 0
  }
}
```

## Post /orders

Adds a new order to the order book.

### Request Body

```json
{
  "take": {
    "token": "0x0000000000000000000000000000000000000000",
    "amount": "3000000000000000000"
  },
  "make": {
    "token": "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07",
    "amount": "300000000000000000000"
  },
  "expires": 1514892553,
  "nonce": 12,
  "maker": "0x9f612fcb422d1971c1be7416c37e3ebc77c0de19",
  "exchange": "0x58e91b0734e2b33efc86067ce4db128366f30dc9",
  "signature": {
    "v": 12,
    "r": "0x9242685bf161793cc25603c231bc2f568eb630ea16aa137d2664ac8038825608",
    "s": "0x4f8ae3bd7535248d0bd448298cc2e2071e56992d0774dc340c368ae950852ada",
    "sig_mode": 0
  }
}
```

#### Sig Modes

| Value | Meaning                                                                                                         |
|-------|-----------------------------------------------------------------------------------------------------------------|
| 0     | Messages signed using the ```eth_signTypedData``` method                                                        |
| 1     | Messages signed using geth, this will append the ```\x19Ethereum Signed Message:\n32``` prefix when verifying   |
| 2     | Messages signed using a trezor, this will append the ```\x19Ethereum Signed Message:\n\x20``` prefix when verifying |
