import { useNavigate } from "react-router-dom";
import BottomNav from "../../components/user/BottomNav";
import "./Dashboard.css";

function Dashboard() {
const navigate = useNavigate();

return ( <div className="dashboard">

```
  {/* HEADER */}
  <div className="navbar">

    <div className="navbar-left">

      <div className="logo-circle">
        🚚
      </div>

      <div>
        <h3>LogiTrack</h3>
        <p>Fast Delivery Service</p>
      </div>

    </div>

    <div className="navbar-right">

      <button className="nav-btn">
        🔔
      </button>

      <button
        className="nav-btn"
        onClick={() => navigate("/profile")}
      >
        👤
      </button>

    </div>

  </div>

  {/* TRACKING SEARCH */}
  <div className="search-box">

    <input
      type="text"
      placeholder="🔍 Track Waybill"
    />

  </div>

  {/* HERO BANNER */}
  <div className="hero-banner">

    <h2>Fast & Reliable Logistics</h2>

    <p>
      Create orders and track shipments easily
    </p>

  </div>

  {/* QUICK ACTION */}
  <h3 className="section-title">
    Quick Actions
  </h3>

  <div className="action-grid">

    <div
      className="action-card"
      onClick={() => navigate("/user/orders/create")}
    >
      <div className="icon">📦</div>

      <span>Create Order</span>
    </div>


    <div
      className="action-card"
      onClick={() => navigate("/tracking-user")}
    >
      <div className="icon">📦</div>
      <span>Track</span>
    </div>

  </div>

  {/* SHIPPING SERVICES */}
  <h3 className="section-title">
    Shipping Services
  </h3>

  <div className="service-grid">

    <div
      className="service-card"
      onClick={() => navigate("/user/orders/create")}
    >
      <h2>EZ</h2>
      <p>Regular</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/user/orders/create")}
    >
      <h2>DOC</h2>
      <p>Document</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/user/orders/create")}
    >
      <h2>JSD</h2>
      <p>Same Day</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/user/orders/create")}
    >
      <h2>JND</h2>
      <p>Next Day</p>
    </div>

  </div>

  <BottomNav />

</div>


);
}

export default Dashboard;
