import { useNavigate, useLocation } from "react-router-dom";

function BottomNav() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div className="bottom-nav">

      <div
        className={location.pathname === "/admin" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin")}
      >
        🏠
        <span>Home</span>
      </div>

      <div
        className={location.pathname === "/admin/orders" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/orders")}
      >
        📦
        <span>Orders</span>
      </div>

      <div
        className={location.pathname === "/admin/pickups" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/pickups")}
      >
        🛎️
        <span>Pickup</span>
      </div>

      <div
        className={location.pathname === "/admin/warehouse" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/warehouse")}
      >
        🏭
        <span>Warehouse</span>
      </div>

      <div
        className={location.pathname === "/admin/shipment" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/shipment")}
      >
        🚛
        <span>Shipment</span>
      </div>

      <div
        className={location.pathname === "/admin/delivery" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/delivery")}
      >
        📬
        <span>Delivery</span>
      </div>

      <div
        className={location.pathname === "/admin/tracking" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/tracking")}
      >
        🚚
        <span>Tracking</span>
      </div>

      <div
        className={location.pathname === "/admin/profile" ? "nav-item active" : "nav-item"}
        onClick={() => navigate("/admin/profile")}
      >
        👤
        <span>Profile</span>
      </div>

    </div>
  );
}

export default BottomNav;
