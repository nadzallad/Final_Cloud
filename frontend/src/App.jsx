import { Routes, Route } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import Orders from "./pages/Orders";
import CreateOrder from "./pages/CreateOrder";
import Tracking from "./pages/Tracking";
import Profile from "./pages/Profile";
import Payment from "./pages/Payment";
import Pickups from "./pages/Pickups";
import Warehouse from "./pages/Warehouse";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/orders" element={<Orders />} />
      <Route path="/orders/create" element={<CreateOrder />} />
      <Route path="/tracking" element={<Tracking />} />
      <Route path="/profile" element={<Profile />} />
      <Route path="/payment" element={<Payment />}/>
      <Route path="/pickups" element={<Pickups />}/>
      <Route path="/warehouse" element={<Warehouse />}/>
    </Routes>
  );
}

export default App;