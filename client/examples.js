const plants = require('./plants/client');

const client = plants.Client('http://localhost:8080');

client.listCriticalPlants()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Critical plants:');
        list.forEach((c) => console.log(c.id + "\t" + c.soilMoistureLevel + "\t" + c.soilDataTimestamp));
    })
    .catch((e) => {
        console.log(`Problem listing critical plants: ${e.message}`);
    });

client.addMoistureLevel(135, 0.1)
    .then(() => {
        console.log('=== Scenario 2 ===');
        console.log('New moisture level record was added');
    })
    .catch((e) => {
        console.log(`Problem adding a moisture level: ${e.message}`);
    });
