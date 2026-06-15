import axios from "axios";

const trackingApi = axios.create({
  baseURL: "http://localhost:8087",
});

const getTrackings = () => {
  return trackingApi.get("/tracking");
};

const createTracking = (data) => {
  return trackingApi.post("/tracking", data);
};

export default {
  getTrackings,
  createTracking,
};
