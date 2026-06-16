import { Routes, Route } from "react-router-dom";
import NotificationListener from "./components/NotificationListener";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

import Dashboard from "./pages/Dashboard";
import Orders from "./pages/Orders";
import CreateOrder from "./pages/CreateOrder";
import Tracking from "./pages/Tracking";
import TrackingUser from "./pages/TrackingUser";
import Profile from "./pages/Profile";
import Payment from "./pages/Payment";
import Pickups from "./pages/Pickups";
import Warehouse from "./pages/Warehouse";
import Shipment from "./pages/Shipment";

function App() {
  <>
    <NotificationListener />
    <ToastContainer />
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/orders" element={<Orders />} />
      <Route path="/orders/create" element={<CreateOrder />} />
      <Route path="/tracking" element={<Tracking />} />
      <Route path="/tracking-user" element={<TrackingUser />} />
      <Route path="/profile" element={<Profile />} />
      <Route path="/payment" element={<Payment />}/>
      <Route path="/pickups" element={<Pickups />}/>
      <Route path="/warehouse" element={<Warehouse />}/>
      <Route path="/shipment" element={<Shipment />} />
    </Routes>
  <?
  );
}

export default App;
