use dexy;

db.orders.createIndex({ expires: 1 }, {expireAfterSeconds: 0});
db.orders.createIndex({ price: 1 });
db.orders.createIndex({ "get.token": 1 });
db.orders.createIndex({ "give.token": 1 });

db.history.createIndex({ timestamp: 1 });