import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { createPayment } from "../services/paymentService";

function Payment() {

  const navigate = useNavigate();

  const orderDraft = JSON.parse(
    localStorage.getItem("orderDraft")
  );

  const [loading, setLoading] =
    useState(false);

  const handlePay = async () => {

    try {

      setLoading(true);

      const paymentData = {
        order_id: Number(
          orderDraft.order_id
        ),
        payment_method:
          orderDraft.payment_method,
        total: Number(
          orderDraft.shipping_cost || 0
        ),
      };

      const res =
        await createPayment(
          paymentData
        );

      if (!res.data?.payment_url) {

        alert(
          "Gagal mendapatkan URL pembayaran"
        );

        return;
      }

      localStorage.removeItem(
        "orderDraft"
      );

      window.location.href =
        res.data.payment_url;

    } catch (err) {

      console.error(err);

      alert(
        err.response?.data?.error ||
        "Gagal membuat pembayaran"
      );

    } finally {

      setLoading(false);

    }
  };
  return (
    <div
      style={{
        maxWidth: "700px",
        margin: "30px auto",
        padding: "20px",
        border: "1px solid #ddd",
        borderRadius: "10px",
      }}
    >

      <h1>Payment</h1>

      <hr />

      <h3>
        Order Information
      </h3>

      <p>
        <b>Sender:</b>{" "}
        {orderDraft.sender_name}
      </p>

      <p>
        <b>Receiver:</b>{" "}
        {orderDraft.receiver_name}
      </p>

      <p>
        <b>Item:</b>{" "}
        {orderDraft.item_name}
      </p>

      <p>
        <b>Weight:</b>{" "}
        {orderDraft.weight_kg} Kg
      </p>

      <p>
        <b>
          Total:
        </b>{" "}
        Rp
        {Number(
          orderDraft.shipping_cost ||
          0
        ).toLocaleString()}
      </p>

      <button
        onClick={handlePay}
        disabled={loading}
        style={{
          marginTop: "20px",
          padding:
            "10px 20px",
        }}
      >
        {loading
          ? "Processing..."
          : "Bayar Sekarang"}
      </button>

      <button
        onClick={() =>
          navigate(
            "/orders/create"
          )
        }
        style={{
          marginLeft: "10px",
          padding:
            "10px 20px",
        }}
      >
        Kembali
      </button>

    </div>
  );
}

export default Payment;