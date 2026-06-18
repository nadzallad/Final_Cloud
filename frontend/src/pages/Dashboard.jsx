import { useNavigate } from "react-router-dom";
import BottomNav from "../components/BottomNav";
import "./Dashboard.css";

function Dashboard() {
const navigate = useNavigate();

return ( <div className="dashboard">

```
  {/* HEADER */}
  <div className="navbar">

    <div className="navbar-left">

      <div className="logo-circle">
        📦
      </div>

      <div>
        <h2 className="brand-name">PaketBang!</h2>
        <p className="brand-tagline">
          Kirim Cepat, Sampai Tepat 
        </p>
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


  {/* HERO BANNER */}
  <div className="hero-banner">

    <div className="hero-content">

      <span className="hero-badge">
        🚀 Delivery Service #1
      </span>

      <h1>
        PaketBang!
      </h1>

      <h2>
        Kirim Cepat,<br />
        Sampai Tepat
      </h2>

      <p>
        Solusi pengiriman modern untuk kebutuhan bisnis dan pribadi.
        Buat pesanan dan lacak paket dengan mudah, cepat, dan aman.
      </p>

      <button
        className="hero-btn"
        onClick={() => navigate("/orders/create")}
      >
        📦 Kirim Paket Sekarang
      </button>

    </div>

  </div>

  {/* QUICK ACTION */}
  <h3 className="section-title">
    Quick Actions
  </h3>

  <div className="action-grid">

    <div
      className="action-card"
      onClick={() => navigate("/orders/create")}
    >
      <div className="icon">📦</div>

      <span>Create Order</span>
    </div>

    <div
      className="action-card"
      onClick={() => navigate("/tracking")}
    >
      <div className="icon">🚚</div>

      <span>Track Parcel</span>
    </div>

    <div
      className="action-card"
      onClick={() => navigate("/tracking-user")}
    >
      <div className="icon">📦</div>
      <span>Track User</span>
    </div>

  </div>

  {/* SHIPPING SERVICES */}
  <h3 className="section-title">
    Shipping Services
  </h3>

  <div className="service-grid">

    <div
      className="service-card"
      onClick={() => navigate("/orders/create")}
    >
      <h2>EZ</h2>
      <p>Regular</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/orders/create")}
    >
      <h2>DOC</h2>
      <p>Document</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/orders/create")}
    >
      <h2>JSD</h2>
      <p>Same Day</p>
    </div>

    <div
      className="service-card"
      onClick={() => navigate("/orders/create")}
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
