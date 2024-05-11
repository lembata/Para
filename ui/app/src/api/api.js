//import fetch from 'node-fetch';
//import { promises as fs } from 'fs';

//TODO: Load from env file
const url = 'http://localhost:8080/';

const API = {
  Accounts: {
    add: async (data) => {
      await requester._post('api/accounts/add', data)
    },
  }
}

const requester = {

  _get: async (slug, data) => {
    return await fetch(`${url}${slug}`)
      .then(response => response.json())
      .catch(error => console.error('Error:', error))
  },

  _post: async (slug, data) => {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    }

    return await fetch(`${url}${slug}`, requestOptions)
      .then(response => response.json())
      .catch(error => console.error('Error:', error))
  }
}

export default API;
