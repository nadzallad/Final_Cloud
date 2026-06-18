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
import Delivery from "./pages/admin/Delivery";

import UserDashboard from "./pages/user/Dashboard";
import UserOrder from "./pages/user/Orders";
import UserCreateOrder from "./pages/user/CreateOrder";
import UserProfile from "./pages/user/Profile";
import TrackingUser from "./pages/user/TrackingUser";

import CourierDashboard from "./pages/courier/Dashboard";
import CourierProfile from "./pages/courier/Profile";
import CourierPickup from "./pages/courier/Pickup";
import CourierDelivery from "./pages/courier/Delivery";
import CourierTracking from "./pages/courier/Tracking";

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
      <Route path="/admin/delivery" element={<Delivery />} />

      {/* User */}
      <Route path="/user" element={<UserDashboard />} />
      <Route path="/user/order" element={<UserOrder />} />
      <Route path="/user/order/create" element={<UserCreateOrder />} />
      <Route path="/user/profile" element={<UserProfile />}/>
      <Route path="/user/tracking" element={<TrackingUser/>} />

      {/* Courier */}
      <Route path="/courier" element={<CourierDashboard />} />
      <Route path="/courier/profile" element={<CourierProfile />} />
      <Route path="/courier/pickup" element={<CourierPickup />} />
      <Route path="/courier/delivery" element={<CourierDelivery />} />
      <Route path="/courier/tracking" element={<CourierTracking />} />
    </Routes>
  );
}

export default App;
