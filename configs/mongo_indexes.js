use OrderBook;

db.Orders.createIndex({ hash: 1 }, {unique: true});
db.Orders.createIndex({ expires: 1 }, {expireAfterSeconds: 0});
db.Orders.createIndex({ price: 1 });
db.Orders.createIndex({ "get.token": 1 });
db.Orders.createIndex({ "give.token": 1 });
