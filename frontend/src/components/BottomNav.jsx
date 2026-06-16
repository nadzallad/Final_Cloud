import { useNavigate, useLocation } from "react-router-dom";

function BottomNav() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div className="bottom-nav">

      <div
        className={
          location.pathname === "/"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/")}
      >
        🏠
        <span>Home</span>
      </div>

      <div
        className={
          location.pathname === "/orders"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/orders")}
      >
        📦
        <span>Orders</span>
      </div>

      <div
        className={
          location.pathname === "/pickups"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/pickups")}
      >
        🛎️
        <span>Pickup</span>
      </div>

      <div
        className={
          location.pathname === "/warehouse"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/warehouse")}
      >
        🏭
        <span>Warehouse</span>
      </div>

      <div
        className={
          location.pathname === "/shipment"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/shipment")}
      >
        🚛
        <span>Shipment</span>
      </div>

      <div
        className={
          location.pathname === "/tracking"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/tracking")}
      >
        🚚
        <span>Tracking</span>
      </div>

      <div
        className={
          location.pathname === "/profile"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/profile")}
      >
        👤
        <span>Profile</span>
      </div>

    </div>
  );
}

export default BottomNav;
