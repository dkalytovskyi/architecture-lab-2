const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listCriticalPlants: () => client.get('/plants'),
        addMoistureLevel: (id, soilMoistureLevel) => client.post('/plants', { id, soilMoistureLevel })
    }

};

module.exports = { Client };
