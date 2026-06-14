import api from "./api";

export const getOrders = () => {
  return api.get("/api/orders");
};

export const createOrder = (data) => {
  return api.post("/api/orders", data);
};