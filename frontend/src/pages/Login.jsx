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
        "http://localhost:5001/auth/login",
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
    <div>
      <h2>Login</h2>

      <form onSubmit={handleSubmit}>

        <div>
          <input
            type="email"
            name="email"
            placeholder="Masukkan Email"
            value={form.email}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <input
            type="password"
            name="password"
            placeholder="Masukkan Password"
            value={form.password}
            onChange={handleChange}
            required
          />
        </div>

        <button type="submit">
          Login
        </button>

      </form>

      <p>
        Belum punya akun?
        <a href="/register"> Register</a>
      </p>
    </div>
  );
}

export default Login;