import { Routes, Route } from "react-router-dom";

import Login from "./pages/Login";
import Register from "./pages/Register";

import AdminDashboard from "./pages/admin/Dashboard";
import Orders from "./pages/admin/Orders";
import CreateOrder from "./pages/admin/CreateOrder";
import Payment from "./pages/admin/Payment";
import Tracking from "./pages/admin/Tracking";
import Warehouse from "./pages/admin/Warehouse";
import Shipment from "./pages/admin/Shipment";
import Profile from "./pages/admin/Profile";
import Pickups from "./pages/admin/Pickups";

import UserDashboard from "./pages/user/Dashboard";
import UserProfile from "./pages/user/Profile";
import TrackingUser from "./pages/user/TrackingUser";

import CourierDashboard from "./pages/courier/Dashboard";
import CourierProfile from "./pages/courier/Profile";

function App() {
  return (
    <Routes>
      {/* Auth */}
      <Route path="/" element={<Login />} />
      <Route path="/register" element={<Register />} />

      {/* Admin */}
      <Route path="/admin" element={<AdminDashboard />} />
      <Route path="/admin/orders" element={<Orders />} />
      <Route path="/admin/orders/create" element={<CreateOrder />} />
      <Route path="/admin/payment" element={<Payment />} />
      <Route path="/admin/tracking" element={<Tracking />} />
      <Route path="/admin/warehouse" element={<Warehouse />} />
      <Route path="/admin/shipment" element={<Shipment />} />
      <Route path="/admin/profile" element={<Profile />} />
      <Route path="/admin/pickups" element={<Pickups />} />

      {/* User */}
      <Route path="/user" element={<UserDashboard />} />
      <Route path="/user/orders" element={<Orders />} />
      <Route path="/user/orders/create" element={<CreateOrder />} />
      <Route path="/user/profile" element={<UserProfile />}/>
      <Route path="/user/tracking" element={<TrackingUser/>} />

      {/* Courier */}
      <Route path="/courier" element={<CourierDashboard />} />
      <Route path="/courier/profile" element={<CourierProfile />} />
    </Routes>
  );
}

export default App;