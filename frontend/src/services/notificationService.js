import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8088",
});

const getNotifications = () => {
  return api.get("/notification");
};

export default {
  getNotifications,
};