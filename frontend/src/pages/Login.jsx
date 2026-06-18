import { useState, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import axios from "axios";

function Login() {
  const { login } = useContext(AuthContext);

  const [form, setForm] = useState({
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
      const response = await axios.post(
        "http://20.249.145.91:8080/auth/login",
        form
      );

      console.log("FULL RESPONSE:", response.data);

      const { token, user } = response.data;

      console.log("TOKEN:", token);
      console.log("USER:", user);
      console.log("ROLE:", user.role);

      login(token, user.role);

      alert(
        `Login berhasil sebagai ${user.role}`
      );

      if (user.role === "admin") {
        window.location.href = "/admin";
      } else if (user.role === "kurir") {
        window.location.href = "/courier";
      } else {
        window.location.href = "/user";
      }

    } catch (error) {

      console.error(error);

      alert(
        error.response?.data?.message ||
        "Email atau password salah"
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
          Login untuk melanjutkan
        </p>

        <form onSubmit={handleSubmit}>
          <div style={{ marginBottom: "15px" }}>
            <input
              type="email"
              name="email"
              placeholder="Masukkan Email"
              value={form.email}
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
              placeholder="Masukkan Password"
              value={form.password}
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
              transition: "0.3s",
              boxShadow: "0 4px 15px rgba(220,38,38,0.3)",
            }}
          >
            Login
          </button>
        </form>

        <p
          style={{
            textAlign: "center",
            marginTop: "20px",
            color: "#666",
          }}
        >
          Belum punya akun?
          <a
            href="/register"
            style={{
              color: "#dc2626",
              fontWeight: "bold",
              textDecoration: "none",
              marginLeft: "5px",
            }}
          >
            Register
          </a>
        </p>
      </div>
    </div>
  );
}

export default Login;