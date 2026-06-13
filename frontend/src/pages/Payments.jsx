import { useEffect, useState } from "react";
import { getPayments } from "../services/paymentApi";
import Navbar from "../components/Navbar";

function Payments() {
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    const res = await getPayments();
    setPayments(res.data);
  };

  return (
    <>
      <Navbar />

      <div style={{ padding: "20px" }}>
        <h1>Payments</h1>

        <table border="1" cellPadding="10">
          <thead>
            <tr>
              <th>Payment ID</th>
              <th>Order ID</th>
              <th>Method</th>
              <th>Total</th>
              <th>Status</th>
            </tr>
          </thead>

          <tbody>
            {payments.map((p) => (
              <tr key={p.payment_id}>
                <td>{p.payment_id}</td>
                <td>{p.order_id}</td>
                <td>{p.payment_method}</td>
                <td>{p.total}</td>
                <td>{p.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}

export default Payments;