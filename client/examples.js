const plants = require('./plants/client');

const client = plants.Client('http://localhost:8080');

// Scenario 1: Set moisture level for a plant.
client.setMoistureLevel(367, 0.8)
    .then((resp) => {
        console.log('=== Scenario 1 ===');
        console.log('Add moisture level response:', resp);
    })
    .catch((e) => {
        console.log(`Problem adding a moisture level: ${e.message}`);
    });

// Scenario 2: Display critical plants.
client.listCriticalPlants()
    .then((list) => {
        console.log('=== Scenario 2 ===');
        console.log('Critical plants:');
        list.forEach((c) => console.log(`Id: ${c.id}; Moisture Level: ${c.soilMoistureLevel}`));
    })
    .catch((e) => {
        console.log(`Problem displaying plants: ${e.message}`);
    });
