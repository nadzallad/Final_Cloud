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
  const [orders, setOrders] = useState([]);
  const [payments, setPayments] = useState([]);
  const [pickups, setPickups] = useState([]);
  const [warehouseLogs, setWarehouseLogs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadOverview();
  }, []);

  const loadOverview = async () => {
    try {
      const res = await warehouseApi.get("/warehouse/overview");
      const data = res.data || {};
      setOrders(data.orders || []);
      setPayments(data.payments || []);
      setPickups(data.pickups || []);
      setWarehouseLogs(data.warehouse_logs || []);
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
      loadOverview();
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
          Ringkasan data dari seluruh tahap pengiriman (order, payment, pickup,
          dan gudang).
        </p>

        <h2 style={sectionTitleStyle}>📦 Orders</h2>
        {orders.length === 0 ? (
          <p>Belum ada data order.</p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table border="1" cellPadding="10" style={tableStyle}>
              <thead style={theadStyle}>
                <tr>
                  <th>Order ID</th>
                  <th>Pengirim</th>
                  <th>Penerima</th>
                  <th>Barang</th>
                  <th>Total</th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {orders.map((o, i) => (
                  <tr key={o.order_id ?? i}>
                    <td>{o.order_id}</td>
                    <td>{o.sender_name}</td>
                    <td>{o.receiver_name}</td>
                    <td>{o.item_name}</td>
                    <td>{o.total_price}</td>
                    <td>
                      <StatusBadge
                        value={o.status}
                        successValues={["DELIVERED", "PAID"]}
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        <h2 style={sectionTitleStyle}>💳 Payments</h2>
        {payments.length === 0 ? (
          <p>
            Belum ada data payment (endpoint list payment mungkin belum
            tersedia di payment-service).
          </p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table border="1" cellPadding="10" style={tableStyle}>
              <thead style={theadStyle}>
                <tr>
                  <th>Payment ID</th>
                  <th>Order ID</th>
                  <th>Metode</th>
                  <th>Total</th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {payments.map((p, i) => (
                  <tr key={p.payment_id ?? i}>
                    <td>{p.payment_id}</td>
                    <td>{p.order_id}</td>
                    <td>{p.payment_method}</td>
                    <td>{p.total}</td>
                    <td>
                      <StatusBadge value={p.status} successValues={["PAID"]} />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        <h2 style={sectionTitleStyle}>🛎️ Pickups</h2>
        {pickups.length === 0 ? (
          <p>Belum ada data pickup.</p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table border="1" cellPadding="10" style={tableStyle}>
              <thead style={theadStyle}>
                <tr>
                  <th>Pickup ID</th>
                  <th>No Resi</th>
                  <th>User ID</th>
                  <th>Berat</th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {pickups.map((p, i) => (
                  <tr key={p.pickup_id ?? i}>
                    <td>{p.pickup_id}</td>
                    <td>{p.tracking_number}</td>
                    <td>{p.user_id}</td>
                    <td>{p.weight_kg} kg</td>
                    <td>
                      <StatusBadge
                        value={p.status}
                        successValues={["PICKED_UP"]}
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

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
