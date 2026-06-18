import { useNavigate, useLocation } from "react-router-dom";

function BottomNav() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div className="bottom-nav">

      <div
        className={
          location.pathname === "/courier"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/courier")}
      >
        🏠
        <span>Home</span>
      </div>

      <div
        className={
          location.pathname === "/courier/pickup"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/courier/pickup")}
      >
        🛎️
        <span>Pickup</span>
      </div>

      <div
        className={
          location.pathname === "/courier/delivery"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/courier/delivery")}
      >
        🚚
        <span>Delivery</span>
      </div>

      <div
        className={
          location.pathname === "/courier/tracking"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/courier/tracking")}
      >
        🚚
        <span>Tracking</span>
      </div>
      
      <div
        className={
          location.pathname === "/courier/profile"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/courier/profile")}
      >
        👤
        <span>Profile</span>
      </div>

    </div>
  );
}

export default BottomNav;
