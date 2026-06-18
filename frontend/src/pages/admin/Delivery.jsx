import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/admin/BottomNav";
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

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>
        <h1>Delivery</h1>

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
                <th>Tracking ID</th>
                <th>Alamat</th>
                <th>Kurir</th>
                <th>Status</th>
                <th>Diterima</th>
                <th>Dibuat</th>
              </tr>
            </thead>
            <tbody>
              {deliveries.map((d) => (
                <tr key={d.delivery_id}>
                  <td>{d.delivery_id}</td>
                  <td>{d.no_resi || "-"}</td>
                  <td>{d.tracking_id}</td>
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
                  <td>{d.created_at ? new Date(d.created_at).toLocaleString("id-ID") : "-"}</td>
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
