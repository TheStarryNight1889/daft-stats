const axios = require('axios').default;

axios.defaults.baseURL = 'http://localhost:3000/query';

// graph query
const propertiesQuery = `
    query {
        properties {
                price
        }
    }
`;
const statsQuery = `
    query {
        stats {
          timestamp
          price_average
          price_high
          price_low
          price_distribution
          properties_added
          properties_removed
        }
    }
`;

const GetAllProperties = async () => axios
  .post('', { query: propertiesQuery })
  .then((response) => response.data)
  .catch((error) => {
    console.log(error);
  });
const GetAllStats = async () => axios
  .post('', { query: statsQuery })
  .then((response) => response.data)
  .catch((error) => {
    console.log(error);
  });

export default {
  GetAllProperties,
  GetAllStats,
};
