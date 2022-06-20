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

const GetAllProperties = async () => axios.post('', { query: propertiesQuery })
  .then((response) => response.data).catch((error) => {
    console.log(error);
  });

export default {
  GetAllProperties,
};
