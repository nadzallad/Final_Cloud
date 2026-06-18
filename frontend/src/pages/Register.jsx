import { useState } from "react";
import axios from "axios";

function Register() {

  const [form, setForm] = useState({
    name: "",
    email: "",
    password: "",
  });

  const handleChange = (e) => {

    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {

    e.preventDefault();

    try {

      await axios.post(
        "http://20.249.145.91:8080/auth/register",
        form
      );

      alert("Register berhasil");

      window.location.href = "/login";

    } catch (err) {

      alert(
        err.response?.data?.message ||
          "Register gagal"
      );
    }
  };

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        background: "linear-gradient(135deg, #991b1b, #dc2626)",
      }}
    >
      <div
        style={{
          width: "100%",
          maxWidth: "400px",
          background: "white",
          padding: "40px",
          borderRadius: "20px",
          borderTop: "5px solid #dc2626",
          boxShadow: "0 10px 30px rgba(0,0,0,0.2)",
        }}
      >
        <h1
          style={{
            textAlign: "center",
            color: "#dc2626",
            marginBottom: "10px",
          }}
        >
          PaketBang!
        </h1>

        <h4
          style={{
            textAlign: "center",
            color: "#991b1b",
            marginBottom: "10px",
          }}
        >
          Kirim cepat, sampai tepat
        </h4>

        <p
          style={{
            textAlign: "center",
            color: "#666",
            marginBottom: "30px",
          }}
        >
          Buat akun baru untuk melanjutkan
        </p>

        <form onSubmit={handleSubmit}>
          <div style={{ marginBottom: "15px" }}>
            <input
              type="text"
              name="name"
              placeholder="Nama Lengkap"
              onChange={handleChange}
              required
              style={{
                width: "100%",
                padding: "14px",
                borderRadius: "10px",
                border: "1px solid #ddd",
                fontSize: "15px",
                outline: "none",
              }}
            />
          </div>

          <div style={{ marginBottom: "15px" }}>
            <input
              type="email"
              name="email"
              placeholder="Email"
              onChange={handleChange}
              required
              style={{
                width: "100%",
                padding: "14px",
                borderRadius: "10px",
                border: "1px solid #ddd",
                fontSize: "15px",
                outline: "none",
              }}
            />
          </div>

          <div style={{ marginBottom: "20px" }}>
            <input
              type="password"
              name="password"
              placeholder="Password"
              onChange={handleChange}
              required
              style={{
                width: "100%",
                padding: "14px",
                borderRadius: "10px",
                border: "1px solid #ddd",
                fontSize: "15px",
                outline: "none",
              }}
            />
          </div>

          <button
            type="submit"
            style={{
              width: "100%",
              padding: "14px",
              background: "#dc2626",
              color: "white",
              border: "none",
              borderRadius: "10px",
              fontSize: "16px",
              fontWeight: "bold",
              cursor: "pointer",
              boxShadow: "0 4px 15px rgba(220,38,38,0.3)",
            }}
          >
            Register
          </button>
        </form>

        <p
          style={{
            textAlign: "center",
            marginTop: "20px",
            color: "#666",
          }}
        >
          Sudah punya akun?
          <a
            href="/"
            style={{
              color: "#dc2626",
              fontWeight: "bold",
              textDecoration: "none",
              marginLeft: "5px",
            }}
          >
            Login
          </a>
        </p>
      </div>
    </div>
  );
}

export default Register;