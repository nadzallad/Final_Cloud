import api from "./api";

export const getOrders = () => {
  return api.get("/orders");
};

export const createOrder = (data) => {
  return api.post("/orders", data);
};

export const confirmPayment = (id) => {
  return api.post(
    `/orders/${id}/confirm-payment`
  );
};