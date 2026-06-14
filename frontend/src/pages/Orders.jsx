import { useEffect, useState } from "react";
import { getOrders } from "../services/orderService";
import Navbar from "../components/Navbar";

function Orders() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    loadOrders();
  }, []);

  const loadOrders = async () => {
    try {
      const res = await orderService.getOrders();
      setOrders(res.data);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <>
      <Navbar />

      <div style={{ padding: "20px" }}>
        <h1>Orders</h1>

        <table border="1" cellPadding="10">
          <thead>
            <tr>
              <th>ID</th>
              <th>Sender</th>
              <th>Receiver</th>
              <th>Item</th>
              <th>Weight</th>
              <th>Origin</th>
              <th>Destination</th>
              <th>Total</th>
              <th>Status</th>
            </tr>
          </thead>

          <tbody>
            {orders.map((o) => (
              <tr key={o.order_id}>
                <td>{o.order_id}</td>
                <td>{o.sender_name}</td>
                <td>{o.receiver_name}</td>
                <td>{o.item_name}</td>
                <td>{o.weight_kg}</td>
                <td>{o.origin_city}</td>
                <td>{o.destination_city}</td>
                <td>{o.total_price}</td>
                <td>{o.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}

export default Orders;