import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function Profile() {
  const navigate = useNavigate();

  const [user, setUser] = useState({
    id: "",
    name: "",
    email: "",
    role: "",
  });

  useEffect(() => {
    const storedUser = JSON.parse(
      localStorage.getItem("user")
    );

    if (storedUser) {
      setUser(storedUser);
    }
  }, []);

  const handleLogout = () => {
    localStorage.clear();
    alert("Logout berhasil");
    window.location.href = "/";
  };

  const getRoleColor = () => {
    if (user.role === "admin") {
      return "#e74c3c";
    }

    if (user.role === "kurir") {
      return "#3498db";
    }

    return "#27ae60";
  };

  return (
    <div
      style={{
        minHeight: "100vh",
        background: "#f5f6fa",
        padding: "30px",
      }}
    >
      <div
        style={{
          maxWidth: "600px",
          margin: "0 auto",
          background: "white",
          borderRadius: "12px",
          padding: "30px",
          boxShadow:
            "0 2px 10px rgba(0,0,0,0.1)",
        }}
      >
        <div
          style={{
            textAlign: "center",
            marginBottom: "30px",
          }}
        >
          <div
            style={{
              width: "100px",
              height: "100px",
              borderRadius: "50%",
              background: "#c0392b",
              color: "white",
              fontSize: "40px",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              margin: "0 auto",
            }}
          >
            👤
          </div>

          <h2>{user.name}</h2>

          <span
            style={{
              background: getRoleColor(),
              color: "white",
              padding: "6px 12px",
              borderRadius: "20px",
              fontSize: "14px",
            }}
          >
            {user.role}
          </span>
        </div>

        <div
          style={{
            display: "flex",
            flexDirection: "column",
            gap: "15px",
          }}
        >
          <div>
            <strong>ID User</strong>
            <p>{user.id || "-"}</p>
          </div>

          <div>
            <strong>Nama</strong>
            <p>{user.name || "-"}</p>
          </div>

          <div>
            <strong>Email</strong>
            <p>{user.email || "-"}</p>
          </div>

          <div>
            <strong>Role</strong>
            <p>{user.role || "-"}</p>
          </div>
        </div>

        <hr
          style={{
            margin: "25px 0",
          }}
        />

        <div
          style={{
            display: "flex",
            gap: "10px",
          }}
        >
          <button
            onClick={() => navigate(-1)}
            style={{
              flex: 1,
              padding: "12px",
              border: "none",
              borderRadius: "8px",
              cursor: "pointer",
            }}
          >
            Kembali
          </button>

          <button
            onClick={handleLogout}
            style={{
              flex: 1,
              padding: "12px",
              background: "#e74c3c",
              color: "white",
              border: "none",
              borderRadius: "8px",
              cursor: "pointer",
            }}
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  );
}

export default Profile;