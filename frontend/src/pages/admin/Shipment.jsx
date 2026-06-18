import { useState } from "react";
import "./Tracking.css";
import { updateShipmentStatus } from "../../services/shipment.js";

function Badge({ status }) {
  let color = "#6b7280";

  if (status === "DELIVERED") color = "#16a34a";
  else if (status === "ON DELIVERY") color = "#2563eb";
  else if (status === "PROCESS") color = "#f59e0b";

  return (
    <span
      style={{
        background: color,
        color: "white",
        padding: "4px 10px",
        borderRadius: 20,
        fontSize: 12,
      }}
    >
      {status}
    </span>
  );
}

function Shipment() {
  const [formData, setFormData] = useState({
    no_resi: "",
    status: "",
    location: "",
    note: "",
  });

  const [shipments, setShipments] = useState([]);
  const [searched] = useState("");

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      await updateShipmentStatus(formData.no_resi, {
        status: formData.status,
        current_location: formData.location,
        note: formData.note,
      });

      alert("Shipment berhasil diperbarui!");

      setFormData({
        no_resi: "",
        status: "",
        location: "",
        note: "",
      });
    } catch (error) {
      alert("Gagal memperbarui shipment");
      console.error(error);
    }
  };

  return (
    <div className="tracking-container">
      {/* Form */}
      <div className="tracking-card">
        <h1>🚚 Shipment Paket</h1>
        <p>Input informasi perjalanan paket</p>

        <form onSubmit={handleSubmit}>
          <input
            type="text"
            name="no_resi"
            placeholder="Nomor Resi"
            value={formData.no_resi}
            onChange={handleChange}
          />

          <input
            type="text"
            name="status"
            placeholder="Status Paket"
            value={formData.status}
            onChange={handleChange}
          />

          <input
            type="text"
            name="location"
            placeholder="Lokasi Saat Ini"
            value={formData.location}
            onChange={handleChange}
          />

          <textarea
            name="note"
            placeholder="Catatan"
            rows="4"
            value={formData.note}
            onChange={handleChange}
          />

          <button type="submit">Update Shipment</button>
        </form>
      </div>

      {/* Tabel */}
      <div
        className="tracking-card"
        style={{ marginTop: 20 }}
      >
        <h2
          style={{
            fontSize: 16,
            marginBottom: 12,
            color: "#1a1a2e",
          }}
        >
          Semua Shipment
        </h2>

        {shipments.length === 0 ? (
          <p style={{ color: "#888" }}>Belum ada shipment.</p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table
              style={{
                width: "100%",
                borderCollapse: "collapse",
                fontSize: 13,
                border: "1px solid #e5e7eb",
              }}
            >
              <thead
                style={{
                  background: "#c0392b",
                  color: "white",
                }}
              >
                <tr>
                  {[
                    "ID",
                    "Tracking ID",
                    "No Resi",
                    "Asal",
                    "Tujuan",
                    "Lokasi Saat Ini",
                    "Status",
                    "ETA",
                  ].map((h) => (
                    <th
                      key={h}
                      style={{
                        padding: "10px 12px",
                        textAlign: "left",
                        fontWeight: 600,
                      }}
                    >
                      {h}
                    </th>
                  ))}
                </tr>
              </thead>

              <tbody>
                {shipments.map((s) => (
                  <tr
                    key={s.shipment_id}
                    style={{
                      background:
                        s.no_resi === searched ? "#eff6ff" : "white",
                      borderBottom: "1px solid #f1f5f9",
                    }}
                  >
                    <td style={{ padding: "10px 12px" }}>
                      {s.shipment_id}
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      {s.tracking_id}
                    </td>

                    <td
                      style={{
                        padding: "10px 12px",
                        fontWeight: 600,
                      }}
                    >
                      {s.no_resi}
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      {s.origin_city}
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      {s.destination_city}
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      {s.current_location}
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      <Badge status={s.status} />
                    </td>

                    <td style={{ padding: "10px 12px" }}>
                      {s.eta
                        ? new Date(s.eta).toLocaleString("id-ID")
                        : "-"}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default Shipment;