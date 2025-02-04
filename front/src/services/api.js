import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

export const getContainers = async () => {
    try {
      const response = await axios.get(`${API_URL}/containers`);
      return response.data;
    } catch (error) {
      console.error('Error fetching containers:', error);
      throw error;
    }
  };