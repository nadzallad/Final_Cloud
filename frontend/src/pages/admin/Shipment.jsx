import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/admin/BottomNav";
import { updateShipmentStatus, getAllShipments } from "../../services/shipment.js";
import "./Dashboard.css";

function Badge({ status }) {
  const map = {
    DELIVERED:        { bg: "#16a34a" },
    "ON DELIVERY":    { bg: "#2563eb" },
    PROCESS:          { bg: "#f59e0b" },
    IN_TRANSIT:       { bg: "#d97706" },
    "In Transit":     { bg: "#d97706" },
    Pending:          { bg: "#6b7280" },
  };
  const color = map[status]?.bg ?? "#6b7280";
  return (
    <span style={{
      background: color, color: "white",
      padding: "4px 10px", borderRadius: 20, fontSize: 12,
    }}>
      {status}
    </span>
  );
}

function Shipment() {
  const [formData, setFormData] = useState({ no_resi: "", status: "", location: "", note: "" });
  const [shipments, setShipments] = useState([]);
  const [loading, setLoading]     = useState(true);

  useEffect(() => { fetchShipments(); }, []);

  const fetchShipments = async () => {
    try {
      const data = await getAllShipments();
      setShipments(Array.isArray(data) ? data : data?.data ?? []);
    } catch (err) {
      console.error("Gagal mengambil data shipment:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
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
      setFormData({ no_resi: "", status: "", location: "", note: "" });
      fetchShipments();
    } catch (err) {
      alert("Gagal memperbarui shipment");
      console.error(err);
    }
  };

  return (
    <>
      <Navbar />

      <div style={{ padding: "20px", paddingBottom: 100 }}>

        <h1>Shipment</h1>

        {/* ── Form Update ── */}
        <div className="form-container" style={{ marginBottom: 24 }}>
          <h1>🚚 Shipment Paket</h1>
          <p style={{ color: "#777", marginBottom: 16 }}>Input informasi perjalanan paket</p>

          <form onSubmit={handleSubmit}>
            <input
              type="text" name="no_resi" placeholder="Nomor Resi"
              value={formData.no_resi} onChange={handleChange}
            />
            <input
              type="text" name="status" placeholder="Status Paket"
              value={formData.status} onChange={handleChange}
            />
            <input
              type="text" name="location" placeholder="Lokasi Saat Ini"
              value={formData.location} onChange={handleChange}
            />
            <textarea
              name="note" placeholder="Catatan" rows="4"
              value={formData.note} onChange={handleChange}
            />
            <button type="submit" className="btn-primary">
              Update Shipment
            </button>
          </form>
        </div>

        {/* ── Tabel Semua Shipment ── */}
        <div style={{
          background: "white", borderRadius: 20, padding: 20,
          boxShadow: "0 4px 15px rgba(0,0,0,.08)",
        }}>
          <h2 style={{ fontSize: 16, marginBottom: 16, color: "#1a1a2e", fontWeight: 700 }}>
            Semua Shipment
          </h2>

          {loading ? (
            <p style={{ color: "#888" }}>Memuat data...</p>
          ) : shipments.length === 0 ? (
            <p style={{ color: "#888" }}>Belum ada shipment.</p>
          ) : (
            <div style={{ overflowX: "auto" }}>
              <table border="1" cellPadding="10"
                style={{ width: "100%", borderCollapse: "collapse" }}>
                <thead style={{ backgroundColor: "#c0392b", color: "white" }}>
                  <tr>
                    {["ID","Tracking ID","No Resi","Asal","Tujuan","Lokasi Saat Ini","Status","ETA"].map(h => (
                      <th key={h}>{h}</th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {shipments.map((s) => (
                    <tr key={s.shipment_id}>
                      <td>{s.shipment_id}</td>
                      <td>{s.tracking_id}</td>
                      <td style={{ fontWeight: 600 }}>{s.no_resi}</td>
                      <td>{s.origin_city}</td>
                      <td>{s.destination_city}</td>
                      <td>{s.current_location}</td>
                      <td><Badge status={s.status} /></td>
                      <td>{s.eta ? new Date(s.eta).toLocaleString("id-ID") : "-"}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>

      </div>

      <BottomNav />
    </>
  );
}

export default Shipment;
