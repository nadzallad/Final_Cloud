import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/courier/BottomNav";
import deliveryApi from "../../services/deliveryService";

const STATUS_COLOR = {
  OUT_FOR_DELIVERY: "#e67e22",
  DELIVERED: "#27ae60",
  FAILED: "#c0392b",
};

const STATUS_LABEL = {
  OUT_FOR_DELIVERY: "Dalam Pengiriman",
  DELIVERED: "Terkirim",
  FAILED: "Gagal Kirim",
};

function Delivery() {
  const [deliveries, setDeliveries] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchResi, setSearchResi] = useState("");
  const [searchResult, setSearchResult] = useState(null);
  const [searchError, setSearchError] = useState("");

  useEffect(() => {
    loadDeliveries();
  }, []);

  const loadDeliveries = async () => {
    try {
      const res = await deliveryApi.get("/deliveries");
      setDeliveries(res.data || []);
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateStatus = async (deliveryID, newStatus) => {
    const label = STATUS_LABEL[newStatus] || newStatus;
    if (!window.confirm(`Tandai sebagai "${label}"?`)) return;
    try {
      await deliveryApi.patch(`/deliveries/${deliveryID}/status`, {
        status: newStatus,
      });
      alert(`Status berhasil diubah ke ${label}!`);
      loadDeliveries();
      if (searchResult?.delivery_id === deliveryID) {
        setSearchResult((prev) => ({ ...prev, status: newStatus }));
      }
    } catch (err) {
      alert("Gagal update status delivery");
    }
  };

  const handleSearch = async () => {
    if (!searchResi.trim()) return;
    setSearchError("");
    setSearchResult(null);
    try {
      const res = await deliveryApi.get(`/deliveries/by-resi/${searchResi.trim()}`);
      setSearchResult(res.data);
    } catch (err) {
      setSearchError("Delivery tidak ditemukan untuk resi tersebut.");
    }
  };

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>
        <h1>Delivery</h1>

        {/* Search by Resi */}
        <div style={{ marginBottom: "20px" }}>
          <div style={{ display: "flex", gap: "8px", marginBottom: "10px" }}>
            <input
              value={searchResi}
              onChange={(e) => setSearchResi(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && handleSearch()}
              placeholder="Cari by no resi..."
              style={{
                padding: "8px 12px",
                border: "1px solid #ddd",
                borderRadius: "4px",
                fontSize: "14px",
                width: "300px",
              }}
            />
            <button
              onClick={handleSearch}
              style={{
                padding: "8px 16px",
                backgroundColor: "#c0392b",
                color: "white",
                border: "none",
                borderRadius: "4px",
                cursor: "pointer",
              }}
            >
              Cari
            </button>
          </div>

          {searchError && <p style={{ color: "#c0392b" }}>{searchError}</p>}

          {searchResult && (
            <div style={{
              padding: "12px",
              border: "1px solid #c0392b",
              borderRadius: "4px",
              marginBottom: "10px",
              backgroundColor: "#fff5f5",
            }}>
              <p><b>Delivery ID:</b> {searchResult.delivery_id}</p>
              <p><b>No Resi:</b> {searchResult.no_resi || "-"}</p>
              <p><b>Alamat:</b> {searchResult.delivery_address}</p>
              <p>
                <b>Status:</b>{" "}
                <span style={{
                  padding: "3px 8px",
                  borderRadius: "4px",
                  backgroundColor: STATUS_COLOR[searchResult.status] || "#95a5a6",
                  color: "white",
                  fontSize: "12px",
                }}>
                  {STATUS_LABEL[searchResult.status] || searchResult.status}
                </span>
              </p>
              {searchResult.delivered_at && (
                <p><b>Diterima:</b> {new Date(searchResult.delivered_at).toLocaleString("id-ID")}</p>
              )}
              {searchResult.status === "OUT_FOR_DELIVERY" && (
                <div style={{ marginTop: "8px", display: "flex", gap: "8px" }}>
                  <button
                    onClick={() => handleUpdateStatus(searchResult.delivery_id, "DELIVERED")}
                    style={{ padding: "4px 10px", backgroundColor: "#27ae60", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" }}
                  >
                    ✓ Tandai Terkirim
                  </button>
                  <button
                    onClick={() => handleUpdateStatus(searchResult.delivery_id, "FAILED")}
                    style={{ padding: "4px 10px", backgroundColor: "#c0392b", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" }}
                  >
                    ✕ Gagal Kirim
                  </button>
                </div>
              )}
            </div>
          )}
        </div>

        {/* Tabel */}
        {loading ? (
          <p>Memuat data...</p>
        ) : deliveries.length === 0 ? (
          <p>Belum ada delivery.</p>
        ) : (
          <table border="1" cellPadding="10" style={{ width: "100%", borderCollapse: "collapse" }}>
            <thead style={{ backgroundColor: "#c0392b", color: "white" }}>
              <tr>
                <th>ID</th>
                <th>No Resi</th>
                <th>Alamat</th>
                <th>Kurir</th>
                <th>Status</th>
                <th>Diterima</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              {deliveries.map((d) => (
                <tr key={d.delivery_id}>
                  <td>{d.delivery_id}</td>
                  <td>{d.no_resi || "-"}</td>
                  <td>{d.delivery_address}</td>
                  <td>
                    {d.courier_name || "-"}
                    {d.courier_phone && (
                      <div style={{ fontSize: "11px", color: "#888" }}>{d.courier_phone}</div>
                    )}
                  </td>
                  <td>
                    <span style={{
                      padding: "4px 8px",
                      borderRadius: "4px",
                      backgroundColor: STATUS_COLOR[d.status] || "#95a5a6",
                      color: "white",
                      fontSize: "12px",
                    }}>
                      {STATUS_LABEL[d.status] || d.status}
                    </span>
                  </td>
                  <td>{d.delivered_at ? new Date(d.delivered_at).toLocaleString("id-ID") : "-"}</td>
                  <td>
                    {d.status === "OUT_FOR_DELIVERY" && (
                      <>
                        <button
                          onClick={() => handleUpdateStatus(d.delivery_id, "DELIVERED")}
                          style={{ padding: "4px 8px", backgroundColor: "#27ae60", color: "white", border: "none", borderRadius: "4px", cursor: "pointer", marginRight: "4px" }}
                        >
                          ✓ Terkirim
                        </button>
                        <button
                          onClick={() => handleUpdateStatus(d.delivery_id, "FAILED")}
                          style={{ padding: "4px 8px", backgroundColor: "#c0392b", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" }}
                        >
                          ✕ Gagal
                        </button>
                      </>
                    )}
                    {d.status === "FAILED" && (
                      <button
                        onClick={() => handleUpdateStatus(d.delivery_id, "OUT_FOR_DELIVERY")}
                        style={{ padding: "4px 8px", backgroundColor: "#e67e22", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" }}
                      >
                        ↩ Kirim Ulang
                      </button>
                    )}
                    {d.status === "DELIVERED" && (
                      <span style={{ color: "#27ae60", fontSize: "12px" }}>✓ Selesai</span>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <BottomNav />
    </>
  );
}

export default Delivery;
