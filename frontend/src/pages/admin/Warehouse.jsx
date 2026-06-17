import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/BottomNav";
import warehouseApi from "../../services/warehouseApi";

function Warehouse() {
  const [logs, setLogs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadLogs();
  }, []);

  const loadLogs = async () => {
    try {
      const res = await warehouseApi.get("/warehouse-logs");
      setLogs(res.data || []);
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  const handleMarkOutForShipment = async (warehouseID) => {
    try {
      await warehouseApi.patch(`/warehouse-logs/${warehouseID}/status`, {
        status: "OUT_FOR_SHIPMENT",
      });
      alert("Paket ditandai siap dikirim!");
      loadLogs();
    } catch (err) {
      alert("Gagal update status warehouse");
    }
  };

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>
        <h1>Warehouse</h1>

        {loading ? (
          <p>Memuat data...</p>
        ) : logs.length === 0 ? (
          <p>Belum ada paket di gudang.</p>
        ) : (
          <table
            border="1"
            cellPadding="10"
            style={{ width: "100%", borderCollapse: "collapse" }}
          >
            <thead style={{ backgroundColor: "#c0392b", color: "white" }}>
              <tr>
                <th>ID</th>
                <th>No Resi</th>
                <th>User ID</th>
                <th>Barang</th>
                <th>Stock</th>
                <th>Status</th>
                <th>Dibuat</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              {logs.map((l) => (
                <tr key={l.warehouse_id}>
                  <td>{l.warehouse_id}</td>
                  <td>{l.tracking_number}</td>
                  <td>{l.user_id}</td>
                  <td>{l.item_name || "-"}</td>
                  <td>{l.stock}</td>
                  <td>
                    <span
                      style={{
                        padding: "4px 8px",
                        borderRadius: "4px",
                        backgroundColor:
                          l.status === "OUT_FOR_SHIPMENT"
                            ? "#27ae60"
                            : "#e67e22",
                        color: "white",
                        fontSize: "12px",
                      }}
                    >
                      {l.status}
                    </span>
                  </td>
                  <td>
                    {l.created_at
                      ? new Date(l.created_at).toLocaleString("id-ID")
                      : "-"}
                  </td>
                  <td>
                    {l.status === "IN_WAREHOUSE" && (
                      <button
                        onClick={() => handleMarkOutForShipment(l.warehouse_id)}
                        style={{
                          padding: "4px 8px",
                          backgroundColor: "#c0392b",
                          color: "white",
                          border: "none",
                          borderRadius: "4px",
                          cursor: "pointer",
                        }}
                      >
                        Siap Kirim
                      </button>
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

export default Warehouse;
