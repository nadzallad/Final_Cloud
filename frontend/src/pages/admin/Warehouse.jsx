import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/admin/BottomNav";
import warehouseApi from "../../services/warehouseApi";

const sectionTitleStyle = {
  marginTop: "30px",
  marginBottom: "10px",
};

const tableStyle = {
  width: "100%",
  borderCollapse: "collapse",
  marginBottom: "10px",
};

const theadStyle = {
  backgroundColor: "#c0392b",
  color: "white",
};

function StatusBadge({ value, successValues }) {
  const isSuccess = successValues.includes(value);
  return (
    <span
      style={{
        padding: "4px 8px",
        borderRadius: "4px",
        backgroundColor: isSuccess ? "#27ae60" : "#e67e22",
        color: "white",
        fontSize: "12px",
      }}
    >
      {value || "-"}
    </span>
  );
}

function Warehouse() {
  const [warehouseLogs, setWarehouseLogs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadWarehouseLogs();
}, []);

  const loadWarehouseLogs = async () => {
    try {
      const res = await warehouseApi.get("/warehouse-logs");

      setWarehouseLogs(res.data || []);
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
      loadWarehouseLogs();
    } catch (err) {
      alert("Gagal update status warehouse");
    }
  };

  if (loading) {
    return (
      <>
        <Navbar />
        <div style={{ padding: "20px" }}>
          <h1>Warehouse</h1>
          <p>Memuat data...</p>
        </div>
        <BottomNav />
      </>
    );
  }

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>

        <h1>Warehouse</h1>
        
        <p style={{ color: "#666" }}>
          Data paket yang berada di gudang.
        </p>

        <h2 style={sectionTitleStyle}>🏭 Warehouse Logs</h2>
        {warehouseLogs.length === 0 ? (
          <p>Belum ada paket di gudang.</p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table border="1" cellPadding="10" style={tableStyle}>
              <thead style={theadStyle}>
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
                {warehouseLogs.map((l) => (
                  <tr key={l.warehouse_id}>
                    <td>{l.warehouse_id}</td>
                    <td>{l.tracking_number}</td>
                    <td>{l.user_id}</td>
                    <td>{l.item_name || "-"}</td>
                    <td>{l.stock}</td>
                    <td>
                      <StatusBadge
                        value={l.status}
                        successValues={["OUT_FOR_SHIPMENT"]}
                      />
                    </td>
                    <td>
                      {l.created_at
                        ? new Date(l.created_at).toLocaleString("id-ID")
                        : "-"}
                    </td>
                    <td>
                      {l.status === "IN_WAREHOUSE" && (
                        <button
                          onClick={() =>
                            handleMarkOutForShipment(l.warehouse_id)
                          }
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
          </div>
        )}
      </div>
      <BottomNav />
    </>
  );
}

export default Warehouse;
