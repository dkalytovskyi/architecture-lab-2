const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listCriticalPlants: () => client.get('/plants'),
        setMoistureLevel: (id, level) => client.post('/plants', { id, level })
    }

};

module.exports = { Client };