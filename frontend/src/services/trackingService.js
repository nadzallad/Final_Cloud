import api from "./api";

export const getTrackings = () => {
  return api.get("/tracking");
};

export const createTracking = (data) => {
  return api.post("/tracking", data);
};