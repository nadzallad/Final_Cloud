import api from "./api";

const getTrackings = () => {
  return api.get("/tracking");
};

const createTracking = (data) => {
  return api.post("/tracking", data);
};

export default {
  getTrackings,
  createTracking,
};
