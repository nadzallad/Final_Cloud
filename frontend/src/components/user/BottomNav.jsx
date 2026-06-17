import { useNavigate, useLocation } from "react-router-dom";

function BottomNav() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div className="bottom-nav">

      <div
        className={
          location.pathname === "/user"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() => navigate("/user")}
      >
        🏠
        <span>Home</span>
      </div>

      <div
        className={
          location.pathname === "/user/orders/create"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/orders/create")
        }
      >
        📦
        <span>Order</span>
      </div>

      <div
        className={
          location.pathname === "/user/orders"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/orders")
        }
      >
        📋
        <span>My Orders</span>
      </div>

      <div
        className={
          location.pathname === "/user/tracking"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/tracking")
        }
      >
        🚚
        <span>Tracking</span>
      </div>

      <div
        className={
          location.pathname === "/user/profile"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/profile")
        }
      >
        👤
        <span>Profile</span>
      </div>

    </div>
  );
}

export default BottomNav;