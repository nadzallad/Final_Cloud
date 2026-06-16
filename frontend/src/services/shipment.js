import api from "./api";

export const getAllShipments = async () => {
  const response = await api.get("/shipments");
  return response.data;
};

export const getShipmentByResi = async (noResi) => {
  const response = await api.get(`/shipments/${noResi}`);
  return response.data;
};
