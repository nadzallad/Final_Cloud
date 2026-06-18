import axios from "axios";

const trackingApi = axios.create({
  baseURL: "http://localhost:8087",
});

const getTrackings = () => {
  return trackingApi.get("/tracking");
};

const getTrackingByResi = (noResi) => {
  return trackingApi.get(`/tracking/${noResi}`);
};

export default {
  getTrackings,
  getTrackingByResi,
};