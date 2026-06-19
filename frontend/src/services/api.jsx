import axios from "axios";

const api = axios.create({
  baseURL: "http://20.249.145.91:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;