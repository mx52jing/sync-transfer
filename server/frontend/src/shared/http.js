import axios from "axios";
export const http = axios.create({
  timeout: 5000
});

http.interceptors.response.use(response => {
  return response.data;
})
