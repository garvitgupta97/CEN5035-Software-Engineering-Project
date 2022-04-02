import axios from 'axios';
import store from './store/configureStore';
import { tokenSelector } from './selectors';


const instance = axios.create({
  baseURL: "http://192.168.0.226:8082",
});

instance.interceptors.request.use((config) => {
  const token = tokenSelector(store.getState());
  config.headers.Authorization = token;
  return config;
});

export default instance;
