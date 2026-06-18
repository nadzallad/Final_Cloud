import axios from "axios";

const shipmentApi = axios.create({
  baseURL: "http://localhost:8085/api",
});

export const getAllShipments = async () => {
  const response = await shipmentApi.get("/shipments");
  return response.data;
};

export const getShipmentByResi = async (noResi) => {
  const response = await shipmentApi.get(`/shipments/by-resi/${noResi}`);
  return response.data;
};

export const updateShipmentStatus = async (noResi, data) => {
  const response = await shipmentApi.patch(
    `/shipments/by-resi/${noResi}/status`,
    data
  );

  return response.data;
};