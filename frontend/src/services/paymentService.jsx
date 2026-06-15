import api from "./paymentApi";

export const createPayment = (data) => {
  return api.post("/payments", data);
};