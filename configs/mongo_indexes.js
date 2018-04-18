use dexy;

db.orders.createIndex({ expires: 1 }, {expireAfterSeconds: 0});
db.orders.createIndex({ price: 1 });
db.orders.createIndex({ "take.token": 1 });
db.orders.createIndex({ "make.token": 1 });
db.orders.createIndex({ maker: 1});

db.history.createIndex({ timestamp: 1 });
db.history.createIndex({ block: 1 });
db.history.createIndex({ maker: 1 });
db.history.createIndex({ taker: 1 });
db.history.createIndex({ "take.token": 1 });
db.history.createIndex({ "make.token": 1 });

db.ticks.createIndex({ "pair.quote": 1 });
db.ticks.createIndex({ timestamp: 1 });