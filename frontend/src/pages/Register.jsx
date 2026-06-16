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
        "http://localhost:5001/auth/register",
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
    <div>
      <h2>Register</h2>

      <form onSubmit={handleSubmit}>

        <input
          type="text"
          name="name"
          placeholder="Nama"
          onChange={handleChange}
        />

        <input
          type="email"
          name="email"
          placeholder="Email"
          onChange={handleChange}
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          onChange={handleChange}
        />

        <button type="submit">
          Register
        </button>

      </form>
    </div>
  );
}

export default Register;