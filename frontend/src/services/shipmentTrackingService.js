import axios from "axios";

const warehouseApi = axios.create({ baseURL: "http://localhost:8084/api" });
const deliveryApi  = axios.create({ baseURL: "http://localhost:8086/api" });
const shipmentApi  = axios.create({ baseURL: "http://localhost:8085/api" }); // shipment-service

// Cek apakah paket ada/pernah ada di gudang
export const getWarehouseByResi = async (noResi) => {
  const res = await warehouseApi.get(`/warehouse-logs/by-resi/${noResi}`);
  return res.data; // { found: true/false, data: {...} }
};

// Ambil data shipment (perjalanan antar kota) by no resi
export const getShipmentByResi = async (noResi) => {
  const res = await shipmentApi.get(`/shipments/by-resi/${noResi}`);
  return res.data;
};

// Ambil data delivery (kurir terakhir) by no resi
export const getDeliveryByResi = async (noResi) => {
  const res = await deliveryApi.get(`/deliveries/by-resi/${noResi}`);
  return res.data;
};

// Ambil semua shipment (untuk tabel admin)
export const getAllShipments = async () => {
  const res = await shipmentApi.get("/shipments");
  return res.data;
};
