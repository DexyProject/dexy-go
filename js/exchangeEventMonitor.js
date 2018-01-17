const Exchange = artifacts.require("");


function log(data) {

}

module.exports = function () {

    Exchange.at("0x0").then(function(exchange) {

        var traded = exchange.Traded({fromBlock: blockNumber, toBlock: "latest"}); // @todo blockNumber

        traded.watch(function (err, data) {
            if (err) {
                console.log(err); // @todo
                return // @todo probably fail
            }

            log(data)
        });

        var cancelled = exchange.Cancelled({fromBlock: blockNumber, toBlock: "latest"}); // @todo blockNumber

        cancelled.watch(function (err, data) {
            if (err) {
                console.log(err); // @todo
                return // @todo probably fail
            }

            log(data)
        });

    }).catch(function (err) {
        console.log("Failed: " + err)
    });
};