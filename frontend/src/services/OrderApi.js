import api from "./api";

export const getOrders = () => api.get("/orders");

export const createOrder = (data) =>
  api.post("/orders", data);