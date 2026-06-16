import { useState } from "react";
import { login } from "../services/authService";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const data = await login(email, password);

      localStorage.setItem("token", data.token);
      localStorage.setItem("role", data.user.role);

      alert("Login berhasil");

      if (data.user.role === "admin") {
        window.location.href = "/admin";
      } else if (data.user.role === "kurir") {
        window.location.href = "/courier";
      } else {
        window.location.href = "/user";
      }
    } catch (err) {
      alert("Login gagal");
    }
  };

  return (
    <div>
      <h2>Login</h2>

      <form onSubmit={handleLogin}>
        <input
          type="email"
          placeholder="Email"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />

        <button type="submit">
          Login
        </button>
      </form>
    </div>
  );
}

export default Login;