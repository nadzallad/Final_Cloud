import { useEffect, useState } from "react";
import { getOrders } from "../services/orderService";
import Navbar from "../components/Navbar";
import api from "../services/api";

function Orders() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    loadOrders();
  }, []);

  const loadOrders = async () => {
    try {
      const res = await getOrders();
      const ordersData = res.data;

      // fetch resi untuk setiap order
      const ordersWithResi = await Promise.all(
        ordersData.map(async (order) => {
          try {
            const resiRes = await api.get(`/api/orders/${order.order_id}/resi`);
            return { ...order, no_resi: resiRes.data.no_resi };
          } catch {
            return { ...order, no_resi: "-" };
          }
        })
      );

      setOrders(ordersWithResi);
    } catch (err) {
      console.log(err);
    }
  };

  const handleConfirmPayment = async (orderID) => {
    try {
      await api.post(`/api/orders/${orderID}/confirm-payment`);
      alert("Payment dikonfirmasi! Resi sedang dibuat...");
      loadOrders();
    } catch (err) {
      alert("Gagal konfirmasi payment");
    }
  };

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>
        <h1>Orders</h1>
        {orders.length === 0 ? (
          <p>Belum ada order.</p>
        ) : (
          <table border="1" cellPadding="10" style={{ width: "100%", borderCollapse: "collapse" }}>
            <thead style={{ backgroundColor: "#c0392b", color: "white" }}>
              <tr>
                <th>ID</th>
                <th>Pengirim</th>
                <th>Penerima</th>
                <th>Barang</th>
                <th>Berat</th>
                <th>Asal</th>
                <th>Tujuan</th>
                <th>Jarak</th>
                <th>Ongkir</th>
                <th>No Resi</th>
                <th>Status</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              {orders.map((o) => (
                <tr key={o.order_id}>
                  <td>{o.order_id}</td>
                  <td>{o.sender_name}</td>
                  <td>{o.receiver_name}</td>
                  <td>{o.item_name}</td>
                  <td>{o.weight_kg} kg</td>
                  <td>{o.origin_city}</td>
                  <td>{o.destination_city}</td>
                  <td>{o.distance_km} km</td>
                  <td>Rp {o.shipping_cost?.toLocaleString("id-ID")}</td>
                  <td>{o.no_resi}</td>
                  <td>
                    <span style={{
                      padding: "4px 8px",
                      borderRadius: "4px",
                      backgroundColor: o.status === "PAID" ? "#27ae60" : "#e67e22",
                      color: "white",
                      fontSize: "12px"
                    }}>
                      {o.status}
                    </span>
                  </td>
                  <td>
                    {o.status === "WAITING_PAYMENT" && (
                      <button
                        onClick={() => handleConfirmPayment(o.order_id)}
                        style={{
                          padding: "4px 8px",
                          backgroundColor: "#c0392b",
                          color: "white",
                          border: "none",
                          borderRadius: "4px",
                          cursor: "pointer"
                        }}
                      >
                        Bayar
                      </button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </>
  );
}

export default Orders;