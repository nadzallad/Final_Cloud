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

// Cari tracking berdasarkan no resi
const getTrackingByResi = (noResi) => {
  return trackingApi.get(`/tracking/${noResi}`);
};

export default {
  getTrackings,
  createTracking,
  getTrackingByResi,
};
