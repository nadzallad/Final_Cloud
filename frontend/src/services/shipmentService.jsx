import { useEffect, useState } from "react";
import {
  getAllShipments,
  getShipmentByResi,
} from "../services/shipmentService";

export default function Shipment() {
  const [shipments, setShipments] = useState([]);
  const [resi, setResi] = useState("");
  const [trackingResult, setTrackingResult] = useState(null);

  useEffect(() => {
    fetchShipments();
  }, []);

  const fetchShipments = async () => {
    try {
      const data = await getAllShipments();
      setShipments(data);
    } catch (error) {
      console.error("Gagal mengambil data shipment:", error);
    }
  };

  const handleTrack = async () => {
    if (!resi) {
      alert("Masukkan nomor resi");
      return;
    }

    try {
      const data = await getShipmentByResi(resi);
      setTrackingResult(data);
    } catch (error) {
      alert("Shipment tidak ditemukan");
      setTrackingResult(null);
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>Shipment Service</h1>

      {/* Tracking */}
      <div style={{ marginBottom: "30px" }}>
        <h2>Track Shipment</h2>

        <input
          type="text"
          placeholder="Masukkan No Resi"
          value={resi}
          onChange={(e) => setResi(e.target.value)}
          style={{ marginRight: "10px" }}
        />

        <button onClick={handleTrack}>
          Track
        </button>

        {trackingResult && (
          <div style={{ marginTop: "20px" }}>
            <p><strong>No Resi:</strong> {trackingResult.no_resi}</p>
            <p><strong>Tracking ID:</strong> {trackingResult.tracking_id}</p>
            <p><strong>Origin:</strong> {trackingResult.origin_city}</p>
            <p><strong>Destination:</strong> {trackingResult.destination_city}</p>
            <p><strong>Current Location:</strong> {trackingResult.current_location}</p>
            <p><strong>Status:</strong> {trackingResult.status}</p>
            <p><strong>ETA:</strong> {trackingResult.eta}</p>
          </div>
        )}
      </div>

      {/* List Shipment */}
      <h2>Daftar Shipment</h2>

      <table border="1" cellPadding="10">
        <thead>
          <tr>
            <th>No Resi</th>
            <th>Origin</th>
            <th>Destination</th>
            <th>Current Location</th>
            <th>Status</th>
            <th>ETA</th>
          </tr>
        </thead>

        <tbody>
          {shipments.map((shipment) => (
            <tr key={shipment.shipment_id}>
              <td>{shipment.no_resi}</td>
              <td>{shipment.origin_city}</td>
              <td>{shipment.destination_city}</td>
              <td>{shipment.current_location}</td>
              <td>{shipment.status}</td>
              <td>{shipment.eta}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
