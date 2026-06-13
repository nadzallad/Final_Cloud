import { useEffect, useState } from "react";
import orderService from "../services/orderService";

function Orders() {

  const [orders, setOrders] = useState([]);

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {

      const data =
        await orderService.getOrders();

      setOrders(data);

    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1>Orders</h1>

      {orders.map((order) => (
        <div key={order.order_id}>
          {order.sender_name}
        </div>
      ))}
    </div>
  );
}

export default Orders;