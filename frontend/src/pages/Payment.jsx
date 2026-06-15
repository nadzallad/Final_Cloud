import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { createPayment } from "../services/paymentService";
import { confirmPayment } from "../services/orderService";

function Payment() {

  const navigate = useNavigate();

  const orderDraft = JSON.parse(
    localStorage.getItem("orderDraft")
  );

  const [loading, setLoading] =
    useState(false);

  if (!orderDraft) {

    return (
      <div
        style={{
          padding: "30px",
        }}
      >
        <h2>
          Order tidak ditemukan
        </h2>

        <button
          onClick={() =>
            navigate("/orders/create")
          }
        >
          Buat Order Baru
        </button>
      </div>
    );
  }

  const handlePay = async () => {

    try {

      setLoading(true);

      const paymentData = {
        order_id:
          Number(orderDraft.order_id),

        payment_method:
          orderDraft.payment_method,

        total:
          Number(
            orderDraft.total_price
          ),
      };

      console.log(
        "PAYMENT DATA:",
        paymentData
      );

      const response =
        await createPayment(
          paymentData
        );

      console.log(
        "PAYMENT RESPONSE:",
        response.data
      );

      if (
        response.data.payment_url
      ) {

        window.location.href =
          response.data.payment_url;

        return;
      }

      await confirmPayment(
        paymentData.order_id
      );

      alert(
        "Pembayaran berhasil"
      );

      localStorage.removeItem(
        "orderDraft"
      );

      navigate("/orders");

    } catch (err) {

      console.error(err);

      alert(
        err.response?.data?.error ||
        err.message
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

      <hr />

      <h3>
        Payment Method
      </h3>

      <p>
        {orderDraft.payment_method}
      </p>

      {orderDraft.payment_method ===
        "TRANSFER" && (
        <div
          style={{
            padding: "10px",
            background:
              "#f5f5f5",
          }}
        >
          <h4>
            Virtual Account BCA
          </h4>

          <p>
            123456789012345
          </p>
        </div>
      )}

      {orderDraft.payment_method ===
        "EWALLET" && (
        <div
          style={{
            padding: "10px",
            background:
              "#f5f5f5",
          }}
        >
          <h4>
            Virtual Account BRI
          </h4>

          <p>
            998877665544
          </p>
        </div>
      )}

      {orderDraft.payment_method ===
        "QRIS" && (
        <div
          style={{
            padding: "10px",
            background:
              "#f5f5f5",
          }}
        >
          <h4>
            QRIS
          </h4>

          <img
            src="https://api.qrserver.com/v1/create-qr-code/?size=250x250&data=LOGISTICS-PAYMENT"
            alt="QRIS"
          />
        </div>
      )}

      {orderDraft.payment_method ===
        "COD" && (
        <div
          style={{
            padding: "10px",
            background:
              "#f5f5f5",
          }}
        >
          <h4>
            Cash On Delivery
          </h4>

          <p>
            Pembayaran dilakukan
            saat barang diterima.
          </p>
        </div>
      )}

      <hr />

      <h3>
        Payment Summary
      </h3>

      <p>
        <b>Distance:</b>{" "}
        {orderDraft.distance_km ??
          "-"} Km
      </p>

      <p>
        <b>
          Shipping Cost:
        </b>{" "}
        Rp
        {Number(
          orderDraft.shipping_cost ||
            0
        ).toLocaleString()}
      </p>

      <p>
        <b>Total:</b>{" "}
        Rp
        {Number(
          orderDraft.total_price ||
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