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
          location.pathname === "/user/order/create"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/order/create")
        }
      >
        📦
        <span>Order</span>
      </div>

      <div
        className={
          location.pathname === "/user/order"
            ? "nav-item active"
            : "nav-item"
        }
        onClick={() =>
          navigate("/user/order")
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